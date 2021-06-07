package model

import (
	"fmt"
	"os"
	"time"
)

type Catch struct {
	CatchDirectory string
}

func (c *Catch) GetFileDirectory() string {
	return fmt.Sprintf("%s.catch_log.csv", c.CatchDirectory)
}

func (c *Catch) SaveToLogFile(e error) {
	fmt.Println(c.GetFileDirectory())
	f, err := os.OpenFile(c.GetFileDirectory(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := time.Now().Format(time.RFC3339)
	_, err = f.Write([]byte(fmt.Sprintf("\n%s,%s", t, e.Error())))
	if err != nil {
		panic(err)
	}
}

func (c *Catch) DeleteLogFile() {
	err := os.Remove(c.GetFileDirectory())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
