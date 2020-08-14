package main

import "fmt"

type USB interface {
	start()
	end()
}

type FlashDisk struct {
	name string
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "storage begin")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "storage end")
}

type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "mouse start")
}

func (m Mouse) end() {
	fmt.Println(m.name, "mouse end")
}

func main() {
	var usb USB = FlashDisk{name: "sandisk"}
	usb.start()
	usb.end()

	usb2 := Mouse{name: "logit"}
	usb2.start()
	usb2.end()
}
