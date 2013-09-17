package main

import (
	"log"
	"os"
)

//the purpose of this command is to generate an amount of load into a logfile
func main() {
	out, _ := os.Create("/mnt/logs/test.out")
	logger := log.New(out, "pre", 0)
	m := make([]int, 1000000000)
	for n := range m {
		logger.Println(n)
		logger.Println("hello")
	}

}
