package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	envLevel = "LOGGING_LEVEL"

	separator  = " "
	timeFormat = time.RFC3339Nano
)

var (
	log logger
	buf sync.Pool
)

type logger struct {
	lvl level
	out io.Writer
	mux sync.Mutex
}

func init() {
	lvl := getLevelOrDefault(os.Getenv(envLevel), infoLevelCode)
	log = logger{lvl: lvl, out: os.Stdout}
	buf = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// SetLevel changes log level at runtime
func SetLevel(s string) error {
	lvl, err := getLevel(s)
	if err != nil {
		return err
	}

	log.lvl = lvl
	return nil
}

// Trace writes a message with level trace
func Trace(obj ...interface{}) {
	write(traceLevelCode, fmt.Sprint(obj...))
}

// Tracef writes a message with level trace
func Tracef(msg string, args ...interface{}) {
	write(traceLevelCode, fmt.Sprintf(msg, args...))
}

// Debug writes a message with level debug
func Debug(obj ...interface{}) {
	write(debugLevelCode, fmt.Sprint(obj...))
}

// Debugf writes a message with level debug
func Debugf(msg string, args ...interface{}) {
	write(debugLevelCode, fmt.Sprintf(msg, args...))
}

// Info writes a message with level info
func Info(obj ...interface{}) {
	write(infoLevelCode, fmt.Sprint(obj...))
}

// Infof writes a message with level info
func Infof(msg string, args ...interface{}) {
	write(infoLevelCode, fmt.Sprintf(msg, args...))
}

// Warn writes a message with level warn
func Warn(obj ...interface{}) {
	write(warnLevelCode, fmt.Sprint(obj...))
}

// Warnf writes a message with level warn
func Warnf(msg string, args ...interface{}) {
	write(warnLevelCode, fmt.Sprintf(msg, args...))
}

// Error writes a message with level error
func Error(obj ...interface{}) {
	write(errorLevelCode, fmt.Sprint(obj...))
}

// Errorf writes a message with level error
func Errorf(msg string, args ...interface{}) {
	write(errorLevelCode, fmt.Sprintf(msg, args...))
}

func write(lvl level, msg string) {
	if lvl < log.lvl {
		return
	}

	log.mux.Lock()
	defer log.mux.Unlock()

	if _, err := log.out.Write(format(lvl, msg)); err != nil {
		fmt.Fprintf(os.Stderr, "logger.write failed with error: %v", err)
	}
}

func format(lvl level, msg string) []byte {
	b := buf.Get().(*bytes.Buffer)
	b.Reset()
	defer buf.Put(b)

	b.WriteString(time.Now().Format(timeFormat))

	b.WriteString(separator)
	b.WriteString(lvl.name())

	b.WriteString(separator)
	b.WriteString(msg)

	b.WriteByte('\n')
	return b.Bytes()
}
