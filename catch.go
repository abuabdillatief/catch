package catch

import (
	"fmt"
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
	fmt.Println(strp)
	var i int
	switch printType {
	case TypeError:
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
			if i == 0 {
				fmt.Println(red(d), dir)
				fmt.Println(red(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Println(red(key), val)
		}
	case TypeWarn:
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
			if i == 0 {
				fmt.Println(yellow(d), dir)
				fmt.Println(yellow(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Println(yellow(key), val)
		}
	case TypeInfo:
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
			if i == 0 {
				fmt.Println(blue(d), dir)
				fmt.Println(blue(ei), fmt.Sprintf(`at line: %s`, yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Println(blue(key), val)
		}
	}
	fmt.Printf("\n%s\n", strp)
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
	fmt.Println("__________________")
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(red("\nOriginal error   :\n"), e.Error())
	fmt.Println("\n__________________")
}

func ErrorStr(e string, m string) {
	line, dir := DirectoryFormater(TypeError)
	fmt.Println("__________________")
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(red("\nOriginal error   :\n"), red(e))
	fmt.Println("\n__________________")
}

func Warn(e error, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Println("__________________")
	fmt.Println(yellow("Warning directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(yellow("\nOriginal message :\n"), e.Error())
	fmt.Println("\n__________________")
}

func WarnStr(e string, m string) {
	line, dir := DirectoryFormater(TypeWarn)
	fmt.Println("__________________")
	fmt.Println(yellow("Warning directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, yellow("Warning info     : "), yellow(fmt.Sprintf("%d", line)), yellow(m))
	fmt.Println(yellow("\nOriginal message :\n"), yellow(m))
	fmt.Println("\n__________________")
}

func Inform(e error) {
	fmt.Println("__________________")
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Println(blue("Current directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Println("\n__________________")
}

func InformStr(e string) {
	fmt.Println("__________________")
	line, dir := DirectoryFormater(TypeInfo)
	fmt.Println(blue("Current directory: "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, blue("Info             : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Println("\n__________________")
}

func (c *Catch) Error(e error) {
	fmt.Println("__________________")
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e.Error()))
	fmt.Println("\n__________________")
}

func (c *Catch) ErrorStr(e string) {
	fmt.Println("__________________")
	line, dir := DirectoryFormater(TypeError)
	fmt.Println(red("Error directory  : "), dir)
	fmt.Printf(`%s at line: %s, message: %s`, red("Error info       : "), yellow(fmt.Sprintf("%d", line)), yellow(e))
	fmt.Println("\n__________________")
}
