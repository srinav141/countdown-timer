package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)

}

func init() {
	file, err := os.OpenFile("countdowntimer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}
