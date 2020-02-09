package logger

import (
	"fmt"
	"strings"
)

type level int

const (
	traceLevelName = "TRACE"
	debugLevelName = "DEBUG"
	infoLevelName  = "INFO"
	warnLevelName  = "WARN"
	errorLevelName = "ERROR"
	offLevelName   = "OFF"
)

const (
	traceLevelCode level = iota
	debugLevelCode
	infoLevelCode
	warnLevelCode
	errorLevelCode
	offLevelCode
)

var names = map[level]string{
	traceLevelCode: traceLevelName,
	debugLevelCode: debugLevelName,
	infoLevelCode:  infoLevelName,
	warnLevelCode:  warnLevelName,
	errorLevelCode: errorLevelName,
	offLevelCode:   offLevelName,
}

var codes = map[string]level{
	traceLevelName: traceLevelCode,
	debugLevelName: debugLevelCode,
	infoLevelName:  infoLevelCode,
	warnLevelName:  warnLevelCode,
	errorLevelName: errorLevelCode,
	offLevelName:   offLevelCode,
}

func (l level) name() string {
	return names[l]
}

func getLevelOrDefault(s string, d level) level {
	lvl, err := getLevel(s)
	if err != nil {
		return d
	}
	return lvl
}

func getLevel(s string) (level, error) {
	lvl, ok := codes[strings.ToUpper(s)]
	if !ok {
		return 0, fmt.Errorf("unsupported logging level string '%s'", s)
	}
	return lvl, nil
}
