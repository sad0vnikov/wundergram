package logger

import (
	"github.com/op/go-logging"
)

//Get returns a logger
func Get(name string) *logging.Logger {
	return logging.MustGetLogger(name)
}
