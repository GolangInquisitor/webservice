package loger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

type (
	Logger struct {
		*logrus.Entry
	}

	WriteHook struct {
		Writer    []io.Writer
		LogLevels []logrus.Level
	}
)

func NewLogger() *Logger {
	return &Logger{new()}
}
func (h *WriteHook) Fire(e *logrus.Entry) error {
	line, err := e.String()
	if err != nil {
		return err
	}
	for _, w := range h.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (h *WriteHook) Levels() []logrus.Level {
	return h.LogLevels
}

func new() *logrus.Entry {
	l := logrus.New()

	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		DisableColors: false,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}

	logdir := "logs"

	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		err = os.MkdirAll(logdir, 0777)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		if err != os.ErrExist {
			panic(err)
		}
	}

	l.SetOutput(io.Discard)

	l.AddHook(&WriteHook{
		Writer:    []io.Writer{f, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	return logrus.NewEntry(l)
}
