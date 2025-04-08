package main

import (
	"log"
	"log/syslog"
)

func main() {
	susLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(susLog)
		log.Print("Everything is fine!")
	}
}
