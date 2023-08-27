package quicklog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

type Level int

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "invalid lvl"
	}

}

type Logger struct {
	mu       sync.Mutex
	out      io.Writer
	bufPool  sync.Pool
	level    Level
	tsFormat string
}

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func New() *Logger {
	return &Logger{
		out: os.Stdout,
		bufPool: sync.Pool{
			New: func() any {
				return new(bytes.Buffer)
			},
		},
		level:    InfoLevel,
		tsFormat: time.RFC3339,
	}
}

func (l *Logger) Debug(msg string) {
	l.handleLog(msg, DebugLevel)
}

func (l *Logger) Info(msg string) {
	l.handleLog(msg, InfoLevel)
}

func (l *Logger) Warn(msg string) {
	l.handleLog(msg, WarnLevel)
}

func (l *Logger) Error(msg string) {
	l.handleLog(msg, ErrorLevel)
}

func (l *Logger) Fatal(msg string) {
	l.handleLog(msg, FatalLevel)
}

func (l *Logger) SetLevel(lvl Level) {
	l.level = lvl
}

func (l *Logger) SetWriter(w io.Writer) {
	l.mu.Lock()
	l.out = w
	l.mu.Unlock()
}

func (l *Logger) handleLog(msg string, level Level) {
	if level < l.level {
		return
	}
	bufW := l.bufPool.Get().(*bytes.Buffer)
	bufW.Write([]byte("timestamp="))
	bufW.Write([]byte(time.Now().Format(time.RFC3339)))
	bufW.Write([]byte(" "))

	bufW.Write([]byte("level="))
	bufW.Write([]byte(l.getString(level)))
	bufW.Write([]byte(" "))

	bufW.Write([]byte("message="))
	bufW.Write([]byte(l.getString(msg)))
	bufW.Write([]byte(" "))
	l.mu.Lock()
	io.Copy(l.out, bufW)
	l.mu.Unlock()
	bufW.Reset()
}

func (l *Logger) getString(val any) string {
	switch v := val.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}
