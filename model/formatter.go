package model

import (
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
