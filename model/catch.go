package model

import (
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write([]byte(fmt.Sprintf("\n%s,%s", time.Now().String(), e.Error())))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Catch) DeleteLogFile() {
	err := os.Remove(c.GetFileDirectory())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
