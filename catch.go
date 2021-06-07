package catch

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

type Catch struct {
	CatchDirectory string
}

type PrintType string

const (
	TypeError PrintType = "Error"
	TypeWarn  PrintType = "Warn"
	TypeInfo  PrintType = "Info"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

func NewLog(logFileName string) *Catch {
	err := os.Remove(logFileName)
	if err != nil {
		return nil
	}
	err = os.Mkdir(logFileName, 0755)
	if err != nil {
		return nil
	}
	return &Catch{
		CatchDirectory: logFileName,
	}
}

func DirectoryFormater(printType PrintType) (line int, res string) {
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	switch printType {
	case TypeError:
		s = append(s, red(d))
	case TypeWarn:
		s = append(s, yellow(d))
	case TypeInfo:
		s = append(s, blue(d))
	}
	res = strings.Join(s, "/")
	return
}

func Error(e error, m string) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(red("\nOriginal error   :\n"), red(e.Error()))
	fmt.Println("\n=================")
}

func ErrorStr(e string, m string) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(red("\nOriginal error   :\n"), red(e))
	fmt.Println("\n=================")
}

func Warn(e error, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Println(yellow("Warning directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(yellow("\nOriginal error   :\n"), yellow(e.Error()))
	fmt.Println("\n=================")
}

func WarnStr(e string, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Println(yellow("Warning directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(yellow("\nOriginal error   :\n"), yellow(m))
	fmt.Println("\n=================")
}

func Inform(e error) {
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Println(blue("Current directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Println("\n=================")
}

func InformStr(e string) {
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Println(blue("Current directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Println("\n=================")
}

func (c *Catch) Error(e error) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Println("\n=================")
}

func (c *Catch) ErrorStr(e string) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Println("\n=================")
}
