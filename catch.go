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

func CustomLog(privateLog map[string]string, printType PrintType) {
	line, dir := DirectoryFormater(printType)
	var l string
	for key, _ := range privateLog {
		if len(key) > len(l) {
			l = key
		}
	}

	cd := "Current directory"
	ei := "Error info"
	if len(l) > len(cd) {
		cd += strings.Repeat(" ", len(l)-len(cd))
	}
	if len(cd) > len(ei) {
		ei += strings.Repeat(" ", len(cd)-len(ei))
	}
	if len(l) > len(ei) {
		ei += strings.Repeat(" ", len(l)-len(ei))
	}
	ei += ":  "
	strp := strings.Repeat("_", len(l))
	fmt.Fprintln(os.Stdout, strp)
	var i int
	for key, val := range privateLog {
		d := cd
		if len(key) != len(cd) {
			if len(key) > len(cd) {
				d += strings.Repeat(" ", len(key)-len(cd))

			} else if len(key) < len(cd) {
				key += strings.Repeat(" ", len(cd)-len(key))
			}
		}
		d += ":  "
		key += ":  "
		switch printType {
		case TypeError:
			if i == 0 {

				fmt.Fprintln(os.Stdout, red(d), dir)
				fmt.Fprintln(os.Stdout, red(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, red(key), val)
		case TypeWarn:
			if i == 0 {

				fmt.Fprintln(os.Stdout, yellow(d), dir)
				fmt.Fprintln(os.Stdout, yellow(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, yellow(key), val)
		case TypeInfo:
			if i == 0 {

				fmt.Fprintln(os.Stdout, blue(d), dir)
				fmt.Fprintln(os.Stdout, blue(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, blue(key), val)
		}
	}

	fmt.Fprintf(os.Stdout, "\n%s\n", strp)
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
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Fprintln(os.Stdout, red("\nOriginal error   :\n"), e.Error())
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func ErrorStr(e string, m string) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Fprintln(os.Stdout, red("\nOriginal error   :\n"), red(e))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func Warn(e error, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, yellow("Warning directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Fprintln(os.Stdout, yellow("\nOriginal message :\n"), e.Error())
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func WarnStr(e string, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, yellow("Warning directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Fprintln(os.Stdout, yellow("\nOriginal message :\n"), yellow(m))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func Inform(e error) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Fprintln(os.Stdout, blue("Current directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func InformStr(e string) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Fprintln(os.Stdout, blue("Current directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func (c *Catch) Error(e error) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := DirectoryFormater(TypeError)
	fmt.Fprintln(os.Stdout, red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func (c *Catch) ErrorStr(e string) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := DirectoryFormater(TypeError)
	fmt.Fprintln(os.Stdout, red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Fprintln(os.Stdout, "\n__________________")
}
