package catch

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	print "github.com/abuabdillatief/catch/PrintType"
	"github.com/fatih/color"
	"github.com/fatih/structs"
)

type CatchLogger struct {
	*log.Logger

	//(catch.CatchLogger).Custom is used to store temporary map
	// to log in your terminal.
	// Any map inserted in this key will be deleted
	// right after each print
	Custom map[string]interface{}
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
	if len(s) == 1 {
		s = []string{}
	}
	switch printType {
	case print.TypeError:
		s = append(s, color.New(color.FgRed, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeWarn:
		s = append(s, color.New(color.FgYellow, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeInfo:
		s = append(s, color.New(color.FgBlue, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeNeutral:
		s = append(s, color.New(color.FgWhite, color.Bold).Add(color.Underline).SprintFunc()(d))
	case print.TypeSuccess:
		s = append(s, color.New(color.FgGreen, color.Bold).Add(color.Underline).SprintFunc()(d))
	}
	if  len(strings.Split(strings.Join(s, "/"), "src/")) > 0 {
		str := strings.Split(strings.Join(s, "/"), "src/")
		if len(str) > 0 {
			return str[1]
		}else {
			return str[0]
		}
	}
	return strings.Join(s, "/")
}

// NewLog will create a new CatchLogger instance
// which can then be used to save logs to a .csv file
func NewLog() CatchLogger {
	f, err := os.OpenFile("./catch.log.csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("date,directory,message"))
	if err != nil {
		log.Fatal(err)
	}
	return C
}

// Print will print any value inserted with specified printing style
func Print(typePrint print.PrintType, values ...interface{}) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	var wd, fn, line string

	baseWord := "Info at index"
	logFunc := func(typePrint print.PrintType) func(key1, key2 interface{}) {
		switch typePrint {
		case print.TypeError:
			wd = Red("Working directory: ")
			fn = Red("Function Name    : ")
			line = Red("Location info    : ")

			baseWord = "Error at index"
			return func(key1, key2 interface{}) {
				log.Println(Red(key1), Red(key2))
			}
		case print.TypeWarn:
			wd = Yellow("Working directory: ")
			fn = Yellow("Function Name    : ")
			line = Yellow("Location info    : ")

			return func(key1, key2 interface{}) {
				log.Println(Yellow(key1), Yellow(key2))
			}
		case print.TypeInfo:
			wd = Blue("Working directory: ")
			fn = Blue("Function Name    : ")
			line = Blue("Location info    : ")

			return func(key1, key2 interface{}) {
				log.Println(Blue(key1), Blue(key2))
			}
		case print.TypeNeutral:
			wd = White("Working directory: ")
			fn = White("Function Name    : ")
			line = White("Location info    : ")

			return func(key1, key2 interface{}) {
				log.Println(White(key1), White(key2))
			}
		case print.TypeSuccess:
			wd = Green("Working directory: ")
			fn = Green("Function Name    : ")
			line = Green("Location info    : ")

			return func(key1, key2 interface{}) {
				log.Println(Green(key1), Green(key2))
			}
		}
		return func(key1, key2 interface{}) {
			log.Println(White(key1), White(key2))
		}
	}(typePrint)

	log.Println(wd, DirectoryFormater(inf[0], typePrint))
	log.Println(fn, White(f[len(f)-1]))
	log.Printf(`%s at line: %s`, line, Yellow(inf[1]))

	for i, key := range values {
		str := fmt.Sprintf("%s %d", baseWord, i)
		if len(str) < 17 {
			str += strings.Repeat(" ", 17-len(str))
		}
		str += ": "

		logFunc(str, key)
	}
}

// PrintArray will print any array of string
func PrintArrayString(arr []string) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	var wd, fn, line string

	baseWord := "Index"
	wd = White("Working directory: ")
	fn = White("Function Name    : ")
	line = White("Location info    : ")

	log.Println(wd, DirectoryFormater(inf[0], print.TypeNeutral))
	log.Println(fn, White(f[len(f)-1]))
	log.Printf(`%s at line: %s`, line, Yellow(inf[1]))

	for i, key := range arr {
		str := fmt.Sprintf("%s %d", baseWord, i)
		if len(str) < 17 {
			str += strings.Repeat(" ", 17-len(str))
		}
		str += ": "

		log.Println(White(str), White(key))
	}
}

// PrintArrayInt will print any array of string
func PrintArrayInt(arr []int) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	var wd, fn, line string

	baseWord := "Index"
	wd = White("Working directory: ")
	fn = White("Function Name    : ")
	line = White("Location info    : ")

	log.Println(wd, DirectoryFormater(inf[0], print.TypeNeutral))
	log.Println(fn, White(f[len(f)-1]))
	log.Printf(`%s at line: %s`, line, Yellow(inf[1]))

	for i, key := range arr {
		str := fmt.Sprintf("%s %d", baseWord, i)
		if len(str) < 17 {
			str += strings.Repeat(" ", 17-len(str))
		}
		str += ": "

		log.Println(White(str), White(key))
	}
}

// PrintArrayFloat will print any array of string
func PrintArrayFloat(arr []float64) {
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")

	pc, _, _, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	var wd, fn, line string

	baseWord := "Index"
	wd = White("Working directory: ")
	fn = White("Function Name    : ")
	line = White("Location info    : ")

	log.Println(wd, DirectoryFormater(inf[0], print.TypeNeutral))
	log.Println(fn, White(f[len(f)-1]))
	log.Printf(`%s at line: %s`, line, Yellow(inf[1]))

	for i, key := range arr {
		str := fmt.Sprintf("%s %d", baseWord, i)
		if len(str) < 17 {
			str += strings.Repeat(" ", 17-len(str))
		}
		str += ": "

		log.Println(White(str), White(key))
	}
}

// PrintStruct will print each structs keys and values
// with specific color type
func PrintStructWithType(printType print.PrintType, s interface{}) {
	if !structs.IsStruct(s) {
		Print(print.TypeError, "no struct")
		return
	}
	final := reflect.Indirect(reflect.ValueOf(s)).Interface()
	t := reflect.TypeOf(final)
	m := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		str := fmt.Sprintf("%c", []byte{t.Field(i).Name[0]})
		if strings.ToUpper(str) == str {
			m[t.Field(i).Name] = reflect.ValueOf(final).Field(i).Interface()
		}
	}
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")
	MapPrint(&printType, m, inf)
}

// PrintStruct will print each structs keys and values
func PrintStruct(s interface{}) {
	if !structs.IsStruct(s) {
		Print(print.TypeError, "no struct")
		return
	}
	final := reflect.Indirect(reflect.ValueOf(s)).Interface()
	t := reflect.TypeOf(final)
	m := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		str := fmt.Sprintf("%c", []byte{t.Field(i).Name[0]})
		if strings.ToUpper(str) == str {
			m[t.Field(i).Name] = reflect.ValueOf(final).Field(i).Interface()
		}
	}
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")
	MapPrint(nil, m, inf)
}

// MapPrint will print keys and val inside a map
func MapPrint(printType *print.PrintType, m map[string]interface{}, inf []string) {

	line := inf[1]
	dir := inf[0]

	var l string
	for key := range m {
		if len(key) > len(l) {
			l = key
		}
	}

	cd := "Current directory"
	ei := "Line info"
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
	var logFunc func(key1, key2, key3 interface{}, i int)

	if printType == nil {
		logFunc = func(key1, key2, key3 interface{}, i int) {
			if i == 0 {
				log.Println(White(key3), DirectoryFormater(dir, print.TypeWarn))
				log.Println(White(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
			}
			log.Println(White(key1), key2)
		}
	} else {
		logFunc = func(typePrint print.PrintType) func(key1, key2, key3 interface{}, i int) {
			switch typePrint {
			case print.TypeError:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(Red(key3), DirectoryFormater(dir, print.TypeError))
						log.Println(Red(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(Red(key1), key2)
				}
			case print.TypeWarn:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(Yellow(key3), DirectoryFormater(dir, print.TypeWarn))
						log.Println(Yellow(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(Yellow(key1), key2)
				}
			case print.TypeInfo:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(Blue(key3), DirectoryFormater(dir, print.TypeInfo))
						log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(Blue(key1), key2)
				}
			case print.TypeNeutral:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(White(key3), DirectoryFormater(dir, print.TypeNeutral))
						log.Println(White(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(White(key1), key2)
				}
			case print.TypeSuccess:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(Green(key3), DirectoryFormater(dir, print.TypeSuccess))
						log.Println(Green(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(Green(key1), key2)
				}
			default:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 {
						log.Println(Blue(key3), DirectoryFormater(dir, print.TypeInfo))
						log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
					}
					log.Println(White(key1), key2)
				}
			}
		}(*printType)
	}

	for key, val := range m {
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
		logFunc(key, val, d, i)
		i++
	}
	log.Printf("%s\n", strp)
}

// SaveToLogfile will save your error message
// to your catch.log.csv file
func (c *CatchLogger) SaveToLogFile(e error) {
	if e == nil {
		e = errors.New("no error")
	}
	f, err := os.OpenFile("./catch.log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

// DeleteLogFile will delete your catch.log.csv file
func (c *CatchLogger) DeleteLogFile() {
	err := os.Remove("./catch.log.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

// CustomLog will log every key and value inside catch.Custom
// after each custom logging, map insde catch.Custom will be deleted
/*
	c := catch.NewLog()

	a := map[string]string{
	"test":"ok",
	}

	c.Customlog(print.TypeInfo)
	===================
	2021/07/12 11:54:57 Current directory:  ~/main.go
	2021/07/12 11:54:57 Error info       :   at line: 16
	2021/07/12 11:54:57 tes              :   tes a
	2021/07/12 11:54:57 foo              :   bar
*/
func (c CatchLogger) CustomLog(printType print.PrintType) {
	if len(c.Custom) == 0 {
		log.Println("no maps inserted in catch.Custom")
		return
	}
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")
	MapPrint(&printType, c.Custom, inf)
	for key := range c.Custom {
		delete(c.Custom, key)
	}
}

// MiddlewareLogger will log all keys and values inside your HTTP header
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

//MiddelwareLoggerWithKeys will log specified keys inside your HTTP header
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
