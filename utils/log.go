package utils

import (
	"github.com/sirupsen/logrus"
)

func SetupLogger(server_host, server_port string) {
	log = logrus.WithFields(logrus.Fields{
		"host": server_host,
		"port": server_port,
	})
	log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}

func GetLogger() *logrus.Entry {
	return log
}

var log *logrus.Entry
