package bridge

import "fmt"

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("print request for windwos")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
