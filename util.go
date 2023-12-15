package main

import (
	"fmt"
	"log"
)

func errHandler(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func printer(s []interface{}) {
	for _, item := range s {
		fmt.Printf(">>>>>>>>>>>> %+v\n", item)
	}
}
