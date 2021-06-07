package model

import (
	"runtime"
	"strings"

	"github.com/fatih/color"
)

type PrintType string

const (
	TypeError PrintType = "Error"
	TypeWarn  PrintType = "Warn"
	TypeInfo  PrintType = "Info"
)

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
)

func DirectoryFormater(printType PrintType) (line int, res string) {
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	switch printType {
	case TypeError:
		s = append(s, Red(d))
	case TypeWarn:
		s = append(s, Yellow(d))
	case TypeInfo:
		s = append(s, Blue(d))
	}
	res = strings.Join(s, "/")
	return
}
