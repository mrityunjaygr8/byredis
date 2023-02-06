package utils

import "net"

type ConnState int

const (
	STATE_REQ ConnState = iota
	STATE_RES
	STATE_END
)

type Loop struct {
	events chan<- Event
	errors <-chan error
}

type Event struct {
	State ConnState
	Conn  net.Conn
}

func NewLoop() *Loop {
	return &Loop{}
}

func (l *Loop) AddEvent(conn net.Conn) {
	l.events <- Event{State: STATE_REQ, Conn: conn}
}

func (l *Loop) RunLoop() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			processClient(e.Conn)
		}
	}()
}
func processClient(connection net.Conn) {
	buffer := make([]byte, MAX_MESSAGE_LENGTH+MAX_LENGTH_SECTION_SIZE)
	_, err := connection.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
	}
	_, msg, err := ReadData(buffer)
	log.Println("Received: ", msg)
	resp, err := WriteData("world")
	if err != nil {
		log.Fatal(err)
	}
	_, err = connection.Write(resp)
	connection.Close()
}
