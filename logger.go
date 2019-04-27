package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	logger = &Logger{Level: levelInfoCode, Out: os.Stdout}
	buffer = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

// Logger ..
type Logger struct {
	Level int
	Out   io.Writer
	mutex sync.Mutex
}

// SetLevel ...
func SetLevel(name string) error {
	levelCode, err := getLevelCode(name)
	if err != nil {
		return err
	}
	logger.Level = levelCode
	return nil
}

// SetOutput ...
func SetOutput(out io.Writer) {
	logger.Out = out
}

// Debug ...
func Debug(obj ...interface{}) {
	if logger.Level <= levelDebugCode {
		log(levelDebugName, fmt.Sprint(obj...))
	}
}

// Info ...
func Info(obj ...interface{}) {
	if logger.Level <= levelInfoCode {
		log(levelInfoName, fmt.Sprint(obj...))
	}
}

// Warn ...
func Warn(obj ...interface{}) {
	if logger.Level <= levelWarnCode {
		log(levelWarnName, fmt.Sprint(obj...))
	}
}

// Error ...
func Error(obj ...interface{}) {
	if logger.Level <= levelErrorCode {
		log(levelErrorName, fmt.Sprint(obj...))
	}
}

// Debugf ...
func Debugf(msg string, args ...interface{}) {
	if logger.Level <= levelDebugCode {
		log(levelDebugName, fmt.Sprintf(msg, args...))
	}
}

// Infof ...
func Infof(msg string, args ...interface{}) {
	if logger.Level <= levelInfoCode {
		log(levelInfoName, fmt.Sprintf(msg, args...))
	}
}

// Warnf ...
func Warnf(msg string, args ...interface{}) {
	if logger.Level <= levelWarnCode {
		log(levelWarnName, fmt.Sprintf(msg, args...))
	}
}

// Errorf ...
func Errorf(msg string, args ...interface{}) {
	if logger.Level <= levelErrorCode {
		log(levelErrorName, fmt.Sprintf(msg, args...))
	}
}

func log(lvl string, msg string) {
	line := format(lvl, msg)

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if _, err := logger.Out.Write(line); err != nil {
		fmt.Fprintf(os.Stderr, "logger.log failed: %v", err)
	}
}

func format(lvl string, msg string) []byte {
	b := buffer.Get().(*bytes.Buffer)
	b.Reset()
	defer buffer.Put(b)

	timeStr := time.Now().Format("2006-01-02T15:04:05.000000-0700")
	b.WriteString(timeStr)

	b.WriteByte(' ')
	b.WriteString(lvl)

	b.WriteByte(' ')
	b.WriteString(msg)

	b.WriteByte('\n')
	return b.Bytes()
}
