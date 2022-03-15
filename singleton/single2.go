package singleton

import (
	"fmt"
	"log"
)

type single2 struct {
}

var singleInstance2 *single2

func init() {
	fmt.Println("Creating Single Instance Now")
	singleInstance2 = &single2{}
}

func GetInstance2() *single2 {
	if singleInstance == nil {
		log.Fatal("Single Instance is nil")
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance2
}
