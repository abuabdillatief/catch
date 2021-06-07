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

func (c *Catch) SaveToLogFile(e error) {
	defer f.Close()
	f, err := os.OpenFile(fmt.Sprintf("%s.csv", c.CatchDirectory), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(fmt.Sprintf("\n%s,%s", time.Now().String(), e.Error())))
	if err != nil {
		log.Fatal(err)
	}
}
