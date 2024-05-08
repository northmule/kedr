package utils

import (
	"fmt"
	"log"
)

func CheckError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func InfoError(err interface{}) {
	if err != nil {
		fmt.Println(err)
	}
}
