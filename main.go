package catch

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type CatchLogger struct {
	*log.Logger
	CatchLogDirectory string
}

type PrintType string

//====================================================
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

var C CatchLogger

//====================================================
func DirectoryFormater(dir string, printType PrintType) string {
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
	return strings.Join(s, "/")
}

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

	log.Println("__________________")
	c.ErrLogOut(e, m, DirectoryFormater(inf[0], TypeError), inf[1], TypeError)
	log.Println("__________________")
}

func (c CatchLogger) ErrorStr(e string, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	log.Println("__________________")
	c.StrLogOut(e, m, DirectoryFormater(inf[0], TypeError), inf[1], TypeError)
	log.Println("__________________")
}

func (c CatchLogger) Warn(e error, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	log.Println("__________________")
	c.ErrLogOut(e, m, DirectoryFormater(inf[0], TypeWarn), inf[1], TypeWarn)
	log.Println("__________________")
}

func (c CatchLogger) WarnStr(e string, m string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	log.Println("__________________")
	c.StrLogOut(e, m, DirectoryFormater(inf[0], TypeWarn), inf[1], TypeWarn)
	log.Println("__________________")
}

func (c CatchLogger) Inform(e error) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	log.Println("__________________")
	c.ErrLogOut(e, "", DirectoryFormater(inf[0], TypeInfo), inf[1], TypeInfo)
	log.Println("__________________")
}

func (c CatchLogger) InformStr(e string) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	log.Println("__________________")
	c.StrLogOut(e, "", DirectoryFormater(inf[0], TypeInfo), inf[1], TypeInfo)
	log.Println("__________________")
}

func (c CatchLogger) ErrLogOut(e error, m, dir, line string, typeError PrintType) {
	switch typeError {
	case TypeInfo:
		log.Println(Blue("Working directory: "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s`, Blue("Error info       : "), Yellow(line))

		if len(e.Error()) > len(dir) {
			log.Println(Blue("\nOriginal error   :\n"), e.Error())
		} else {
			log.Println(Blue("\nOriginal error   : "), e.Error())
		}
	case TypeWarn:
		log.Println(Yellow("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, Yellow("Error info       : "), Yellow(line), Yellow(m))

		if len(e.Error()) > len(dir) {
			log.Println(Yellow("\nOriginal error   :\n"), e.Error())
		} else {
			log.Println(Yellow("\nOriginal error   : "), e.Error())
		}
	case TypeError:
		log.Println(Red("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, Red("Error info       : "), Yellow(line), Yellow(m))

		if len(e.Error()) > len(dir) {
			log.Println(Red("\nOriginal error   :\n"), e.Error())
		} else {
			log.Println(Red("\nOriginal error   : "), e.Error())
		}
	}
}

func (c CatchLogger) StrLogOut(e, m, dir, line string, typeError PrintType) {
	switch typeError {
	case TypeInfo:
		log.Println(Blue("Working directory: "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s`, Blue("Error info       : "), Yellow(line))

		if len(e) > len(dir) {
			log.Println(Blue("\nOriginal error   :\n"), e)
		} else {
			log.Println(Blue("\nOriginal error   : "), e)
		}
	case TypeWarn:
		log.Println(Yellow("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, Yellow("Error info       : "), Yellow(line), Yellow(m))

		if len(e) > len(dir) {
			log.Println(Yellow("\nOriginal error   :\n"), e)
		} else {
			log.Println(Yellow("\nOriginal error   : "), e)
		}
	case TypeError:
		log.Println(Red("Error directory  : "), dir)
		fmt.Fprintf(os.Stdout, `%s at line: %s, message: %s`, Red("Error info       : "), Yellow(line), Yellow(m))

		if len(e) > len(dir) {
			log.Println(Red("\nOriginal error   :\n"), e)
		} else {
			log.Println(Red("\nOriginal error   : "), e)
		}
	}
}

func (c *CatchLogger) SaveToLogFile(e error) {
	f, err := os.OpenFile(c.CatchLogDirectory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	err := os.Remove(c.CatchLogDirectory)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (c CatchLogger) CustomLog(privateLog map[string]string, printType PrintType) {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	line := inf[1]
	dir := inf[0]

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

				log.Println(Red(d), DirectoryFormater(dir, TypeError))
				log.Println(Red(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Red(key), val)
		case TypeWarn:
			if i == 0 {

				log.Println(Yellow(d), DirectoryFormater(dir, TypeWarn))
				log.Println(Yellow(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Yellow(key), val)
		case TypeInfo:
			if i == 0 {

				log.Println(Blue(d), DirectoryFormater(dir, TypeInfo))
				log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				i++
			}
			log.Println(Blue(key), val)
		}
	}
	fmt.Fprintf(os.Stdout, "%s\n", strp)
}

func (c CatchLogger) HttpMiddlewareLogger(createLog bool) func(http.Handler) http.Handler {
	var B bytes.Buffer
	c.Logger = log.New(&B, "", log.Llongfile)
	c.Output(2, "")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := make(map[string]string)
			for k, v := range r.Header {
				headers[k] = v[0]
			}
			var l string
			for key, _ := range headers {
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
			fmt.Fprintf(os.Stdout, "%s\n", strp)
			next.ServeHTTP(w, r)
		})
	}
}

func (c CatchLogger) HttpMiddlewareLoggerWithKeys(createLog bool, keys ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := make(map[string]string)
			targetKeys := make(map[string]bool)
			for _, k := range keys {
				targetKeys[k] = true
			}

			for k, v := range r.Header {
				if targetKeys[k] {
					headers[k] = v[0]
				}
			}
			var l string
			for key, _ := range headers {
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
			fmt.Fprintf(os.Stdout, "%s\n", strp)
			next.ServeHTTP(w, r)
		})
	}
}
