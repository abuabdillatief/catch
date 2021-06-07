package catch

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/catch/model"
	"github.com/fatih/color"
)

const (
	TypeError = model.TypeError
	TypeWarn  = model.TypeWarn
	TypeInfo  = model.TypeInfo
)

func NewLog(logFile string) *model.Catch {
	f, err := os.OpenFile(fmt.Sprintf("./%s.catch_log.csv", logFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("date,file,line,message"))
	if err != nil {
		log.Fatal(err)
	}
	return &model.Catch{
		CatchDirectory: logFile,
	}
}

func CustomLog(privateLog map[string]string, printType model.PrintType) {
	fmt.Fprintln(os.Stdout, color.FgYellow, color.FgRed, color.FgBlue)
	line, dir := model.DirectoryFormater(printType)
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
		case model.TypeError:
			if i == 0 {

				fmt.Fprintln(os.Stdout, model.Red(d), dir)
				fmt.Fprintln(os.Stdout, model.Red(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, model.Red(key), val)
		case model.TypeWarn:
			if i == 0 {

				fmt.Fprintln(os.Stdout, model.Yellow(d), dir)
				fmt.Fprintln(os.Stdout, model.Yellow(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, model.Yellow(key), val)
		case model.TypeInfo:
			if i == 0 {

				fmt.Fprintln(os.Stdout, model.Blue(d), dir)
				fmt.Fprintln(os.Stdout, model.Blue(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
				i++
			}
			fmt.Fprintln(os.Stdout, model.Blue(key), val)
		}
	}
	fmt.Fprintf(os.Stdout, "\n%s\n", strp)
}

func Error(e error, m string) {
	line, dir := model.DirectoryFormater(model.TypeError)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, model.Red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Red("Error info       : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(m))
	fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   :\n"), e.Error())
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func ErrorStr(e string, m string) {
	line, dir := model.DirectoryFormater(model.TypeError)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, model.Red("Error directory  : "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Red("Error info       : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(m))
	fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   :\n"), model.Red(e))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func Warn(e error, m string) {
	line, dir := model.DirectoryFormater(model.TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, model.Yellow("Warning directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Yellow("Warning info     : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(m))
	fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal message :\n"), e.Error())
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func WarnStr(e string, m string) {
	line, dir := model.DirectoryFormater(model.TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
	fmt.Fprintln(os.Stdout, model.Yellow("Warning directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Yellow("Warning info     : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(m))
	fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal message :\n"), model.Yellow(m))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func Inform(e error) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := model.DirectoryFormater(model.TypeInfo)
	fmt.Fprintln(os.Stdout, model.Blue("Current directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Blue("Info             : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(e.Error()))
	fmt.Fprintln(os.Stdout, "\n__________________")
}

func InformStr(e string) {
	fmt.Fprintln(os.Stdout, "__________________")
	line, dir := model.DirectoryFormater(model.TypeInfo)
	fmt.Fprintln(os.Stdout, model.Blue("Current directory: "), dir)
	fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Blue("Info             : "), model.Yellow(fmt.Sprintf("%d", line)), model.Yellow(e))
	fmt.Fprintln(os.Stdout, "\n__________________")
}
