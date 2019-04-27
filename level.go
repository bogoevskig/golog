package logger

import (
	"fmt"
	"strings"
)

// ...
const (
	levelDebugCode = 1
	levelDebugName = "DEBUG"
	levelInfoCode  = 2
	levelInfoName  = "INFO"
	levelWarnCode  = 3
	levelWarnName  = "WARN"
	levelErrorCode = 4
	levelErrorName = "ERROR"
	levelOffCode   = 10
	levelOffName   = "OFF"
)

// GetLevelCode returns level code from string or error if string is unknown
func getLevelCode(name string) (int, error) {
	switch strings.ToUpper(name) {
	case levelDebugName:
		return levelDebugCode, nil
	case levelInfoName:
		return levelInfoCode, nil
	case levelWarnName:
		return levelWarnCode, nil
	case levelErrorName:
		return levelErrorCode, nil
	case levelOffName:
		return levelOffCode, nil
	}
	return 0, fmt.Errorf("invalid level name: %s", name)
}
