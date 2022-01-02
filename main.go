package catch

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

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
	Yellow       = color.New(color.FgYellow).SprintFunc()
	Red          = color.New(color.FgRed).SprintFunc()
	Blue         = color.New(color.FgBlue).SprintFunc()
	Green        = color.New(color.FgGreen).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	YellowItalic = color.New(color.FgYellow).Add(color.Italic).SprintFunc()
	RedItalic    = color.New(color.FgRed).Add(color.Italic).SprintFunc()
	BlueItalic   = color.New(color.FgBlue).Add(color.Italic).SprintFunc()
	GreenItalic  = color.New(color.FgGreen).Add(color.Italic).SprintFunc()
	WhiteItalic  = color.New(color.FgWhite).Add(color.Italic).SprintFunc()
)

var C CatchLogger

//====================================================
func directoryFormater(dir string, printType print.PrintType) string {
	arr := strings.Split(dir, "/")
	s := arr[:len(arr)-1]
	d := arr[len(arr)-1]
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
	if len(strings.Split(strings.Join(s, "/"), "src/")) > 1 {
		str := strings.Split(strings.Join(s, "/"), "src/")
		if len(str) > 1 {
			return str[1]
		} else {
			return str[0]
		}
	}
	return strings.Join(s, "/")
}

// Print will print any value inserted with specified printing style
func _print(typePrint print.PrintType, values ...interface{}) {
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

	log.Println(wd, directoryFormater(inf[0], typePrint))
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

// printStruct will print each structs keys and values
func printStruct(s interface{}) {
	if !structs.IsStruct(s) {
		_print(print.TypeError, "no struct")
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
	mapPrint(nil, m, inf, 0)
}

// printStruct will print each structs keys and values
func printStructIndented(s interface{}, indentation int) {
	if !structs.IsStruct(s) {
		_print(print.TypeError, "no struct")
		return
	}
	final := reflect.Indirect(reflect.ValueOf(s)).Interface()
	t := reflect.TypeOf(final)
	m := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		str := fmt.Sprintf("%c", []byte{t.Field(i).Name[0]})
		if strings.ToUpper(str) == str {
			key := ""
			for j := 0; j < indentation; j++ {
				key += "  "
			}
			m[fmt.Sprintf(key+t.Field(i).Name)] = reflect.ValueOf(final).Field(i).Interface()
		}
	}
	var B bytes.Buffer
	clg := log.New(&B, "", log.Llongfile)
	clg.Output(2, "")
	inf := strings.Split(fmt.Sprintf("%v", &B), ":")
	mapPrint(nil, m, inf, indentation)
}

// mapPrint will print keys and val inside a map
func mapPrint(printType *print.PrintType, m map[string]interface{}, inf []string, indentation int) {

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
	// strp := strings.Repeat("_", len(l))
	var i int
	var logFunc func(key1, key2, key3 interface{}, i int)

	if printType == nil {
		logFunc = func(key1, key2, key3 interface{}, i int) {
			if i == 0 && indentation == 0 {
				log.Println("==================")
				log.Println(White(key3), directoryFormater(dir, print.TypeWarn))
				log.Println(White(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
				log.Println("==================")
			}
			if key2 != nil {
				log.Println(White(key1), key2)
			} else {
				log.Println(WhiteItalic(key1))
			}
		}
	} else {
		logFunc = func(typePrint print.PrintType) func(key1, key2, key3 interface{}, i int) {
			switch typePrint {
			case print.TypeError:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(Red(key3), directoryFormater(dir, print.TypeError))
						log.Println(Red(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(Red(key1), key2)
					} else {
						log.Println(RedItalic(key1))

					}
				}
			case print.TypeWarn:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(Yellow(key3), directoryFormater(dir, print.TypeWarn))
						log.Println(Yellow(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(Yellow(key1), key2)
					} else {
						log.Println(YellowItalic(key1))

					}
				}
			case print.TypeInfo:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(Blue(key3), directoryFormater(dir, print.TypeInfo))
						log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(Blue(key1), key2)
					} else {
						log.Println(BlueItalic(key1))

					}
				}
			case print.TypeNeutral:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(White(key3), directoryFormater(dir, print.TypeNeutral))
						log.Println(White(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(White(key1), key2)
					} else {
						log.Println(WhiteItalic(key1))

					}
				}
			case print.TypeSuccess:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(Green(key3), directoryFormater(dir, print.TypeSuccess))
						log.Println(Green(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(Green(key1), key2)
					} else {
						log.Println(GreenItalic(key1))

					}
				}
			default:
				return func(key1, key2, key3 interface{}, i int) {
					if i == 0 && indentation == 0 {
						log.Println("==================")
						log.Println(Blue(key3), directoryFormater(dir, print.TypeInfo))
						log.Println(Blue(ei), fmt.Sprintf(`at line: %s`, Yellow(line)))
						log.Println("==================")
					}

					if key2 != nil {
						log.Println(White(key1), key2)
					} else {
						log.Println(WhiteItalic(key1))

					}
				}
			}
		}(*printType)
	}

	for key, val := range m {
		d := cd
		if reflect.ValueOf(val).Kind() == reflect.Struct {
			logFunc(key, nil, d, i)
			indentation++
			printStructIndented(val, indentation)
		} else {
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
		}
		i++
	}
	if indentation > 0 {
		indentation--
	}
	// log.Printf("%s\n", strp)
}

func Print(s interface{}) {
	switch reflect.ValueOf(s).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(s)
		baseWord := "Index"
		cd := "Current directory"
		ei := "Line info"

		var l string
		for i := 0; i < arr.Len(); i++ {
			if len(arr.Index(i).String()) > len(l) {
				l = arr.Index(i).String()
			}
		}

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

		var B bytes.Buffer
		clg := log.New(&B, "", log.Llongfile)
		clg.Output(2, "")
		inf := strings.Split(fmt.Sprintf("%v", &B), ":")

		log.Println("==================")
		log.Println(White(cd), directoryFormater(inf[0], print.TypeWarn))
		log.Println(White(ei), fmt.Sprintf(`at line: %s`, Yellow(inf[1])))
		log.Println("==================")
		for i := 0; i < arr.Len(); i++ {
			str := fmt.Sprintf("%s %d", baseWord, i)
			if len(str) < 17 {
				str += strings.Repeat(" ", 17-len(str))
			}
			str += ": "
			log.Println(White(str), White(arr.Index(i)))
		}
		fmt.Println("")
	case reflect.Struct:
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
		mapPrint(nil, m, inf, 0)
		fmt.Println("")
	case reflect.Map:
		newMap := make(map[string]interface{})
		v := reflect.ValueOf(s)
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			newMap[key.String()] = strct.Interface()
		}

		var B bytes.Buffer
		clg := log.New(&B, "", log.Llongfile)
		clg.Output(2, "")
		inf := strings.Split(fmt.Sprintf("%v", &B), ":")

		pt := print.TypeNeutral
		mapPrint(&pt, newMap, inf, 0)
		fmt.Println("")
	default:
		fmt.Println(reflect.ValueOf(s).Kind())
	}
}
