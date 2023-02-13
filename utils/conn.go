package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/mrityunjaygr8/byredis/cmd/store"
)

type ConnState int

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "6739"
)

var (
	ErrBadRequest = errors.New("Badly formed request recieved.")
)

const (
	STATE_REQ ConnState = iota
	STATE_RES
	STATE_END
)

type Loop struct {
	events chan<- Event
	errors <-chan error
	store  store.Store
}

type Event struct {
	State ConnState
	Conn  net.Conn
}

func NewLoop() *Loop {
	store := store.NewStore()
	return &Loop{store: store}
}

func (l *Loop) AddEvent(conn net.Conn) {
	l.events <- Event{State: STATE_REQ, Conn: conn}
}

func (l *Loop) RunLoop() {
	events := make(chan Event, 16)
	l.events = events

	errorChan := make(chan error, 1)
	l.errors = errorChan

	go func() {
		for e := range events {
			cmds := e.getCommand()
			switch cmds[0] {
			case SET:
				if len(cmds) != 3 {
					e.respond(ErrBadRequest.Error())
					continue
				}
				l.store.Set(cmds[1], cmds[2])
				log.Info(fmt.Sprintf("Key `%s` set to value `%s` successfully", cmds[1], cmds[2]))

				e.respond(fmt.Sprintf("Key `%s` set to value `%s` successfully", cmds[1], cmds[2]))
			case GET:
				if len(cmds) != 2 {
					e.respond(ErrBadRequest.Error())
					continue
				}
				val, err := l.store.Get(cmds[1])
				if err != nil {
					if errors.Is(err, store.ErrKeyNotFound) {
						sb := strings.Builder{}
						sb.WriteString(err.Error())
						sb.WriteString(fmt.Sprintf(" Missing key: %s", cmds[1]))
						e.respond(sb.String())
						log.Info(sb.String())
						continue
					}
					sb := strings.Builder{}
					sb.WriteString("An error has occurred. ")
					sb.WriteString(err.Error())
					e.respond(sb.String())
					log.Error(err)
					continue
				}

				e.respond(val)

			case DEL:
				if len(cmds) != 2 {
					e.respond(ErrBadRequest.Error())
					continue
				}
				l.store.Del(cmds[1])
				e.respond(fmt.Sprintf("Key %s deleted successfully", cmds[1]))
			}
		}
	}()
}
func (e Event) getCommand() []string {
	buffer := make([]byte, MAX_MESSAGE_LENGTH+MAX_LENGTH_SECTION_SIZE)
	_, err := e.Conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
	}
	return DecodeCommand(buffer)
}

func (e Event) respond(message string) {
	resp, err := WriteData(message)
	if err != nil {
		log.Fatal(err)
	}
	_, err = e.Conn.Write(resp)
	e.Conn.Close()
}
