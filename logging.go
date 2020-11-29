package socketio

import (
	"os"
	"strconv"
)

const (
	db1 = true
)

var (
	DbLogMessage  = true
	LogMessage, _ = strconv.ParseBool(os.Getenv("SOCKETIO_LOG_MESSAGE"))
)
