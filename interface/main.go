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

func checkInt(data interface{}) {
	intData, ok := data.(int)
	if ok {
		number := intData + 10
		fmt.Println("Integer data Type", number)
		return
	}
	float64Data, ok := data.(float64)
	if ok {
		newFloat := float64Data + 10.0
		fmt.Println("Float data Type", newFloat)
		return
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
	checkInt(11.11)
}
