package helper

import (
	"io"
	"time"

	"github.com/qnify/api-server/utils/consts"
	"github.com/qnify/api-server/utils/errors"
	"github.com/zerodha/logf"
)

var Logger logf.Logger
var loggerInit = false

func InitLogger() *logf.Logger {
	if loggerInit {
		panic(errors.New("logger is already initialised"))
	}

	// don't log debug level in production
	level := logf.InfoLevel
	if consts.Dev {
		level = logf.DebugLevel
	}

	Logger = logf.New(logf.Opts{
		Level:                level,
		TimestampFormat:      time.DateTime,
		CallerSkipFrameCount: 3,
		EnableColor:          consts.Dev,
	})

	loggerInit = true
	return &Logger
}

func InitTestLogger(writer io.Writer) *logf.Logger {
	Logger = logf.New(logf.Opts{
		Writer: writer,
	})
	return &Logger
}
