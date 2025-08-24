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

func test(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("Integer:", data)
	case float64:
		fmt.Println("Float64:", data)
	case string:
		fmt.Println("String:", data)
	case bool:
		fmt.Println("Boolean:", data)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var m Device = phone{}
	m.turnOn()
	m.turnOff()

	test(42)
	test("Hello")
	test(true)
	test(3.14)
}
