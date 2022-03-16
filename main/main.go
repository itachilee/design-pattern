package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/itachilee/designpattern/builder"
	"github.com/itachilee/designpattern/factory"
	"github.com/itachilee/designpattern/mediator"
	"github.com/itachilee/designpattern/memento"
	"github.com/itachilee/designpattern/observer"
	"github.com/itachilee/designpattern/prototype"
)

func main() {
	JobTest()
	time.Sleep(1 << 10)
}

func Test() {
	builder.Test()

	fmt.Println()
	factory.Test()
	fmt.Println()
	prototype.Test()
	fmt.Println()

	mediator.Test()
	fmt.Println()

	memento.Test()
	fmt.Println()

	observer.Test()
}

type job struct {
	name string
}

func worker(jobChan <-chan job) {
	for job := range jobChan {
		fmt.Printf("dequeue jobchan %s\n", job.name)
	}
}

func JobTest() {
	jobChan := make(chan job)
	go worker(jobChan)
	for i := 0; i < 100; i++ {
		jobChan <- job{name: strconv.Itoa(i)}
		fmt.Printf("enqueue jobchan %s\n", strconv.Itoa(i))
	}
}
