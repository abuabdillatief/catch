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

func Error(e error) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory  : "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e.Error()))
	fmt.Println("\n=================")
}

func ErrorStr(e string) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory  : "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e))
	fmt.Println("\n=================")
}

func Warn(e error) {
	yellow := color.New(color.FgYellow).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.YellowString("Warning directory: "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e.Error()))
	fmt.Println("\n=================")
}

func WarnStr(e string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.YellowString("Warning directory: "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e))
	fmt.Println("\n=================")
}

func (c *Catch) Error(e error) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory  : "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e.Error()))
	fmt.Println("\n=================")
}

func (c *Catch) ErrorStr(e string) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory  : "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e))
	fmt.Println("\n=================")
}
