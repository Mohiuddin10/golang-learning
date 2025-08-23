package main

import "fmt"

type Device interface {
	turnOn()
	turnOff()
}

type phone struct{}

func (p phone) turnOn() {
	fmt.Println("Phone is on")
}

func (p phone) turnOff() {
	fmt.Println("Phone is off")
}

func main() {
	var m Device = phone{}
	m.turnOn()
	m.turnOff()
}
