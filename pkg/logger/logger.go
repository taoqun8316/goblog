package logger

import (
	"fmt"
	"log"
)

func LogError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Println("error")
	}
}
