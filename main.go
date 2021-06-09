package catch

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/catch/model"
)

const (
	TypeError = model.TypeError
	TypeWarn  = model.TypeWarn
	TypeInfo  = model.TypeInfo
)

type CatchLogger struct {
	*log.Logger
	CatchLogDirectory string
}

var C CatchLogger

func NewLog(logFile string) CatchLogger {
	C.CatchLogDirectory = fmt.Sprintf("./%s.clog.csv", logFile)
	f, err := os.OpenFile(C.CatchLogDirectory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("date,directory,message"))
	if err != nil {
		log.Fatal(err)
	}
	return C
}

func (c CatchLogger) Error(e error, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.ErrLogOut(e, m, DirectoryFormater(inf[0], model.TypeError), inf[1], model.TypeError)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) ErrorStr(e string, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.StrLogOut(e, m, DirectoryFormater(inf[0], model.TypeError), inf[1], model.TypeError)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) Warn(e error, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.ErrLogOut(e, m, DirectoryFormater(inf[0], model.TypeWarn), inf[1], model.TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) WarnStr(e string, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.StrLogOut(e, m, DirectoryFormater(inf[0], model.TypeWarn), inf[1], model.TypeWarn)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) Inform(e error) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.ErrLogOut(e, "", DirectoryFormater(inf[0], model.TypeInfo), inf[1], model.TypeInfo)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) InformStr(e string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	fmt.Fprintln(os.Stdout, "__________________")
	c.StrLogOut(e, "", DirectoryFormater(inf[0], model.TypeInfo), inf[1], model.TypeInfo)
	fmt.Fprintln(os.Stdout, "__________________")
}

func (c CatchLogger) ErrLogOut(e error, m, dir, line string, typeError model.PrintType) {
	switch typeError {
	case model.TypeInfo:
		fmt.Fprintln(os.Stdout, model.Blue("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s`, model.Blue("Error info       : "), model.Yellow(line))

		if len(e.Error()) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Blue("\nOriginal error   :\n"), e.Error())
		} else {
			fmt.Fprintln(os.Stdout, model.Blue("\nOriginal error   : "), e.Error())
		}
	case model.TypeWarn:
		fmt.Fprintln(os.Stdout, model.Yellow("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Yellow("Error info       : "), model.Yellow(line), model.Yellow(m))

		if len(e.Error()) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal error   :\n"), e.Error())
		} else {
			fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal error   : "), e.Error())
		}
	case model.TypeError:
		fmt.Fprintln(os.Stdout, model.Red("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Red("Error info       : "), model.Yellow(line), model.Yellow(m))

		if len(e.Error()) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   :\n"), e.Error())
		} else {
			fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   : "), e.Error())
		}
	}
}

func (c CatchLogger) StrLogOut(e, m, dir, line string, typeError model.PrintType) {
	switch typeError {
	case model.TypeInfo:
		fmt.Fprintln(os.Stdout, model.Blue("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s`, model.Blue("Error info       : "), model.Yellow(line))

		if len(e) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Blue("\nOriginal error   :\n"), e)
		} else {
			fmt.Fprintln(os.Stdout, model.Blue("\nOriginal error   : "), e)
		}
	case model.TypeWarn:
		fmt.Fprintln(os.Stdout, model.Yellow("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Yellow("Error info       : "), model.Yellow(line), model.Yellow(m))

		if len(e) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal error   :\n"), e)
		} else {
			fmt.Fprintln(os.Stdout, model.Yellow("\nOriginal error   : "), e)
		}
	case model.TypeError:
		fmt.Fprintln(os.Stdout, model.Red("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, model.Red("Error info       : "), model.Yellow(line), model.Yellow(m))

		if len(e) > len(dir) {
			fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   :\n"), e)
		} else {
			fmt.Fprintln(os.Stdout, model.Red("\nOriginal error   : "), e)
		}
	}
}

func DirectoryFormater(dir string, printType model.PrintType) string {
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	switch printType {
	case TypeError:
		s = append(s, model.Red(d))
	case TypeWarn:
		s = append(s, model.Yellow(d))
	case TypeInfo:
		s = append(s, model.Blue(d))
	}
	return strings.Join(s, "/")
}

func (c *CatchLogger) GetLogDirectory() string {
	return c.CatchLogDirectory
}

func (c *CatchLogger) SaveToLogFile(e error) {
	f, err := os.OpenFile(c.GetLogDirectory(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	dir := strings.Split(fmt.Sprintf("%v", &B), "/")
	d := dir[len(dir)-1]
	s := d[:len(d)-3]

	t := time.Now().Format(time.RFC3339)
	_, err = f.Write([]byte(fmt.Sprintf("\n%s,%s,%s", t, s, e.Error())))
	if err != nil {
		panic(err)
	}
}

func (c *CatchLogger) DeleteLogFile() {
	err := os.Remove(c.GetLogDirectory())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

// func CustomLog(privateLog map[string]string, printType model.PrintType) {
// 	fmt.Fprintln(os.Stdout, color.FgYellow, color.FgRed, color.FgBlue)
// 	line, dir := model.DirectoryFormater(printType)
// 	var l string
// 	for key, _ := range privateLog {
// 		if len(key) > len(l) {
// 			l = key
// 		}
// 	}

// 	cd := "Current directory"
// 	ei := "Error info"
// 	if len(l) > len(cd) {
// 		cd += strings.Repeat(" ", len(l)-len(cd))
// 	}
// 	if len(cd) > len(ei) {
// 		ei += strings.Repeat(" ", len(cd)-len(ei))
// 	}
// 	if len(l) > len(ei) {
// 		ei += strings.Repeat(" ", len(l)-len(ei))
// 	}
// 	ei += ":  "
// 	strp := strings.Repeat("_", len(l))
// 	fmt.Fprintln(os.Stdout, strp)
// 	var i int
// 	for key, val := range privateLog {
// 		d := cd
// 		if len(key) != len(cd) {
// 			if len(key) > len(cd) {
// 				d += strings.Repeat(" ", len(key)-len(cd))

// 			} else if len(key) < len(cd) {
// 				key += strings.Repeat(" ", len(cd)-len(key))
// 			}
// 		}
// 		d += ":  "
// 		key += ":  "
// 		switch printType {
// 		case model.TypeError:
// 			if i == 0 {

// 				fmt.Fprintln(os.Stdout, model.Red(d), dir)
// 				fmt.Fprintln(os.Stdout, model.Red(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
// 				i++
// 			}
// 			fmt.Fprintln(os.Stdout, model.Red(key), val)
// 		case model.TypeWarn:
// 			if i == 0 {

// 				fmt.Fprintln(os.Stdout, model.Yellow(d), dir)
// 				fmt.Fprintln(os.Stdout, model.Yellow(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
// 				i++
// 			}
// 			fmt.Fprintln(os.Stdout, model.Yellow(key), val)
// 		case model.TypeInfo:
// 			if i == 0 {

// 				fmt.Fprintln(os.Stdout, model.Blue(d), dir)
// 				fmt.Fprintln(os.Stdout, model.Blue(ei), fmt.Sprintf(`at line: %s`, model.Yellow(fmt.Sprintf("%d", line))))
// 				i++
// 			}
// 			fmt.Fprintln(os.Stdout, model.Blue(key), val)
// 		}
// 	}
// 	fmt.Fprintf(os.Stdout, "\n%s\n", strp)
// }
