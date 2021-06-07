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

func NewLog(catchDirectoryPath string) *Catch {
	err := os.Remove(catchDirectoryPath)
	if err != nil {
		return nil
	}
	err = os.Mkdir(catchDirectoryPath, 0755)
	if err != nil {
		return nil
	}
	return &Catch{
		CatchDirectory: catchDirectoryPath,
	}
}

func Error(e error) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory: "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info     : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e.Error()))
}

func (c *Catch) Error(e error) {
	red := color.New(color.FgRed).SprintFunc()
	_, dir, line, _ := runtime.Caller(0)
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	s = append(s, color.RedString(d))

	fmt.Println(color.RedString("Error directory: "), strings.Join(s, "/"))
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info     : "), color.YellowString(fmt.Sprintf("%d", line)), color.YellowString(e.Error()))
}
