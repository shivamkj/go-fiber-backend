package errors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type stack []uintptr

type Err struct {
	message string
	err     error
	stack   *stack
}

func New(message string) error {
	return &Err{
		message: message,
		err:     nil,
		stack:   callers(),
	}
}

func Newf(message string, a ...any) error {
	return &Err{
		message: fmt.Sprintf(message, a...),
		err:     nil,
		stack:   callers(),
	}
}

func Wrap(message string, err error) error {
	return &Err{
		message: message,
		err:     err,
		stack:   callers(),
	}
}

func (err *Err) Error() string {
	return err.message
}

func (err *Err) Stack() string {
	return err.stack.StackTrace()
}

func (s stack) StackTrace() string {
	frames := runtime.CallersFrames(s)
	var message strings.Builder
	for i := 0; i < len(s); i++ {
		frame, _ := frames.Next()
		message.WriteString(frame.Function + "\n\t" + frame.File + ":" + strconv.Itoa(frame.Line) + "\n")
	}
	return message.String()
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	count := runtime.Callers(3, pcs[:])
	var stack stack = pcs[0:count]
	return &stack
}
