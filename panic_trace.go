package utiltools

import (
	"bytes"
	"runtime"
)

// PanicTrace trace panic stack info.
func PanicTrace() string {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, 2048)
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return string(stack)
}
