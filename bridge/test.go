package bridge

func Test() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	mac := &mac{}
	windows := &windows{}
	mac.setPrinter(hpPrinter)
	windows.setPrinter(epsonPrinter)
	mac.print()
	windows.print()
}
