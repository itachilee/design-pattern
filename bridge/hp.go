package bridge

import "fmt"

type hp struct{}

func (p *hp) printFile() {
	fmt.Println("printing by hp printer")
}
