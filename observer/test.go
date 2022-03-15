package observer

func Test() {
	shirtItem := newItem("GoLang Design Patterns")
	observerFirst := &customer{id: "xxx@www.com"}
	observerSecond := &customer{id: "yyy@www.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()

}
