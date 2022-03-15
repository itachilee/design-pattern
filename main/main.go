package main

import (
	"fmt"

	"github.com/itachilee/designpattern/builder"
	"github.com/itachilee/designpattern/factory"
	"github.com/itachilee/designpattern/prototype"
)

func main() {
	builder.Test()
	fmt.Println()
	factory.Test()
	fmt.Println()
	prototype.Test()
	fmt.Println()
}
