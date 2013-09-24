package util

import (
	"log"
)

func LogPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
