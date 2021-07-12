package catch

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	print "github.com/abuabdillatief/catch/PrintType"
	"github.com/fatih/color"
)

type CatchLogger struct {
	*log.Logger
	Custom map[string]string
}

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	White  = color.New(color.FgWhite).SprintFunc()
)

var C CatchLogger

//====================================================
func DirectoryFormater(dir string, printType print.PrintType) string {
	s := strings.Split(dir, "/")
	d := s[len(s)-1]
	s = s[:len(s)-1]
	switch printType {
	case print.TypeError:
		s = append(s, color.New(color.FgRed, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeWarn:
		s = append(s, color.New(color.FgYellow, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeInfo:
		s = append(s, color.New(color.FgBlue, color.Bold).Add(color.Underline).SprintFunc()(d))
	}
	return strings.Join(s, "/")
}

func NewLog() CatchLogger {
	f, err := os.OpenFile("./catch.clog.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("date,directory,message"))
	if err != nil {
		log.Fatal(err)
	}
	return C
}

func Print(typePrint print.PrintType, values ...interface{}) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	baseWord := "Info "
	logFunc := func(typePrint print.PrintType) func(key1, key2 interface{}) {
		switch typePrint {
		case print.TypeError:
			baseWord = "Error"
			return func(key1, key2 interface{}) {
				log.Println(Red(key1), Red(key2))
			}
		case print.TypeWarn:
			return func(key1, key2 interface{}) {
				log.Println(Yellow(key1), Yellow(key2))
			}
		case print.TypeInfo:
			return func(key1, key2 interface{}) {
				log.Println(Blue(key1), Blue(key2))
			}
		case print.TypeNeutral:
			return func(key1, key2 interface{}) {
				log.Println(White(key1), White(key2))
			}
		case print.TypeSuccess:
			return func(key1, key2 interface{}) {
				log.Println(Green(key1), Green(key2))
			}
		}
		return func(key1, key2 interface{}) {
			log.Println(White(key1), White(key2))
		}
	}(typePrint)

	log.Println(Blue("Working directory: "), DirectoryFormater(inf[0], typePrint))
	log.Println(Blue("Function Name    : "), White(f[len(f)-1]))
	log.Printf(`%s at line: %s`, Blue("Location info    : "), Yellow(inf[1]))

	for i, key := range values {
		str := fmt.Sprintf("%s %d", baseWord, i)
		if len(str) < 17 {
			str += strings.Repeat(" ", 17-len(str))
		}
		str += ": "

		logFunc(str, key)
	}
}

func (c CatchLogger) Error(err error, message string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.ErrLogOut(f[len(f)-1], err, message, DirectoryFormater(inf[0], print.TypeError), inf[1], print.TypeError)
	log.Println("__________________")
}

func (c CatchLogger) ErrorStr(err string, message string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.StrLogOut(f[len(f)-1], err, message, DirectoryFormater(inf[0], print.TypeError), inf[1], print.TypeError)
	log.Println("__________________")
}

func (c CatchLogger) Warn(err error, message string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.ErrLogOut(f[len(f)-1], err, message, DirectoryFormater(inf[0], print.TypeWarn), inf[1], print.TypeWarn)
	log.Println("__________________")
}

func (c CatchLogger) WarnStr(err string, message string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.StrLogOut(f[len(f)-1], err, message, DirectoryFormater(inf[0], print.TypeWarn), inf[1], print.TypeWarn)
	log.Println("__________________")
}

func (c CatchLogger) Inform(err error) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.ErrLogOut(f[len(f)-1], err, "", DirectoryFormater(inf[0], print.TypeInfo), inf[1], print.TypeInfo)
	log.Println("__________________")
}

func (c CatchLogger) InformStr(e string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	log.Println("__________________")
	c.StrLogOut(f[len(f)-1], e, "", DirectoryFormater(inf[0], print.TypeInfo), inf[1], print.TypeInfo)
	log.Println("__________________")
}

func (c CatchLogger) ErrLogOut(funcName string, e error, m, dir, line string, typeError print.PrintType) {
	if e == nil {
		e = errors.New("no errors DETECTED")
		typeError = print.TypeSuccess
	}
	switch typeError {
	case print.TypeInfo:
		log.Println(Blue("Working directory: "), dir)
		log.Println(Blue("Function Name    : "), Blue(funcName))

		log.Printf(`%s at line: %s`, Blue("Location info    : "), Yellow(line))

		if len(e.Error()) > len(dir) {
			log.Println(Blue("Original error   :\n"), e.Error())
		} else {
			log.Println(Blue("Original error   : "), e.Error())
		}
	case print.TypeWarn:
		log.Println(Yellow("Error directory  : "), dir)
		log.Println(Yellow("Function Name    : "), Yellow(funcName))
		log.Printf(`%s at line: %s, message: %s`, Yellow("Error info       : "), Yellow(line), Yellow(m))

		if len(e.Error()) > len(dir) {
			log.Println(Yellow("Original error   :\n"), e.Error())
		} else {
			log.Println(Yellow("Original error   : "), e.Error())
		}
	case print.TypeError:
		log.Println(Red("Error directory  : "), dir)
		log.Println(Red("Function Name    : "), Red(funcName))
		log.Printf(`%s at line: %s, message: %s`, Red("Error info       : "), Yellow(line), Yellow(m))

		if len(e.Error()) > len(dir) {
			log.Println(Red("Original error   :\n"), e.Error())
		} else {
			log.Println(Red("Original error   : "), e.Error())
		}
	case print.TypeSuccess:
		log.Println(White("Error directory  : "), dir)
		log.Println(White("Function Name    : "), Green(funcName))
		log.Printf(`%s at line: %s, message: %s`, White("Error info       : "), Green(line), Green(m))

		if len(e.Error()) > len(dir) {
			log.Println(White("Original error   :\n"), e.Error())
		} else {
			log.Println(White("Original error   : "), e.Error())
		}

	}
}

func (c CatchLogger) StrLogOut(funcName, e, m, dir, line string, typeError print.PrintType) {
	switch typeError {
	case print.TypeInfo:
		log.Println(Blue("Working directory: "), dir)
		log.Println(Blue("Function Name    : "), funcName)
		log.Printf(`%s at line: %s`, Blue("Error info       : "), Yellow(line))

		if len(e) > len(dir) {
			log.Println(Blue("Original error   :\n"), e)
		} else {
			log.Println(Blue("Original error   : "), e)
		}
	case print.TypeWarn:
		log.Println(Yellow("Error directory  : "), dir)
		log.Println(Yellow("Function Name    : "), funcName)
		log.Printf(`%s at line: %s, message: %s`, Yellow("Error info       : "), Yellow(line), Yellow(m))

		if len(e) > len(dir) {
			log.Println(Yellow("Original error   :\n"), e)
		} else {
			log.Println(Yellow("Original error   : "), e)
		}
	case print.TypeError:
		log.Println(Red("Error directory  : "), dir)
		log.Println(Red("Function Name    : "), funcName)
		log.Printf(`%s at line: %s, message: %s`, Red("Error info       : "), Yellow(line), Yellow(m))

		if len(e) > len(dir) {
			log.Println(Red("Original error   :\n"), e)
		} else {
			log.Println(Red("Original error   : "), e)
		}
	}
}

func (c *CatchLogger) SaveToLogFile(e error) {
	if e == nil {
		e = errors.New("no error")
	}
	f, err := os.OpenFile("./catch.clog.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
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
	err := os.Remove("./catch.clog.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (c CatchLogger) CustomLog(printType print.PrintType) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	line := inf[1]
	dir := inf[0]

	var l string
	for key := range c.Custom {
		if len(key) > len(l) {
			l = key
		}
	}

	cd := "Current directory"
	ei := "Error info"
	if len(l) > len(cd) {
		cd += strings.Repeat(" ", len(l)-len(cd))
	} else {
		l += strings.Repeat(" ", len(cd)-len(l))
	}

	if len(cd) > len(ei) {
		ei += strings.Repeat(" ", len(cd)-len(ei))
	}
	if len(l) > len(ei) {
		ei += strings.Repeat(" ", len(l)-len(ei))
	}
	ei += ":  "
	strp := strings.Repeat("_", len(l))
	log.Println(strp)
	var i int
	for key, val := range c.Custom {
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
		case print.TypeError:
			if i == 0 {

				log.Println(Red(d), DirectoryFormater(dir, print.TypeError))
				log.Println(Red(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Red(key), val)
		case print.TypeWarn:
			if i == 0 {

				log.Println(Yellow(d), DirectoryFormater(dir, print.TypeWarn))
				log.Println(Yellow(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Yellow(key), val)
		case print.TypeInfo:
			if i == 0 {

				log.Println(Blue(d), DirectoryFormater(dir, print.TypeInfo))
				log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Blue(key), val)
		}
	}
	log.Printf("%s\n", strp)
}

func (c CatchLogger) MiddlewareLogger(createLog bool) func(http.Handler) http.Handler {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := make(map[string]string)
			for k, v := range r.Header {
				headers[k] = v[0]
			}
			var l string
			for key := range headers {
				if len(key) > len(l) {
					l = key
				}
			}
			strp := strings.Repeat("_", len(l))
			log.Println(strp)
			var i int
			for key, val := range headers {
				key += ":  "
				if i == 0 {
					i++
				}
				log.Println(Blue(key), val)
			}
			log.Printf("%s\n", strp)
			next.ServeHTTP(w, r)
		})
	}
}

func (c CatchLogger) MiddlewareLoggerWithKeys(createLog bool, keys ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := make(map[string]string)
			targetKeys := make(map[string]bool)
			for _, k := range keys {
				targetKeys[strings.ToLower(k)] = true
			}

			for k, v := range r.Header {
				if targetKeys[strings.ToLower(k)] {
					headers[k] = v[0]
				}
			}
			var l string
			for key := range headers {
				if len(key) > len(l) {
					l = key
				}
			}

			for key, val := range headers {
				if len(key) < len(l) {
					key += strings.Repeat(" ", len(l)-len(key))
				}
				key += ":  "
				log.Println(Blue(key), val)
			}
			next.ServeHTTP(w, r)
		})
	}
}
