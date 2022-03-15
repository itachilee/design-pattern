package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single3 struct {
}

var singleInstance3 *single3

func GetInstance3() *single3 {
	if singleInstance3 == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now")
				singleInstance3 = &single3{}
			})
		fmt.Println("Single Instance already created-1")
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance3
}
