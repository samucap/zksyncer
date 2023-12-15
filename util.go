package main

import (
	"fmt"
	"log"
)

func errHandler(err error) {
	if err != nil {
		fmt.Printf("Error >>>>>>>>>>>> %+v\n", err)
		panic(err)
	}
}

func printer(s interface{}) {
	log.Printf("printando >>>>>>>>>>>>>>> %v", s)
}
