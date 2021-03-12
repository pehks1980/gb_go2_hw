package main

import (
	"fmt"
	"strconv"
)

type Color int

const (
	Green Color = iota
	Red
	Blue
	Black
)

func (h Human) String() string {
	go PrintStruct()
	return fmt.Sprintf("Greeting: %s, Meters: %d", h.Greeting, h.Meters)
}

func main() {
	h := Human{Greeting: "Hello", Meters: 10}
	s := Stringer(h)

	fmt.Println(ToString(s))        // Greeting: Hello, Meters: 10
	fmt.Println(ToString(107))      // 107
	fmt.Println(ToString("302"))    // 302
	fmt.Println(ToString(3.141592)) // ???
}

// преобразовывание типов на лету
func ToString(any interface{}) string {
	switch v := any.(type) {
	case Stringer:
		//проверяем что any содержит тип у case, в случае успеха его значение v
		return v.String() //возвращаемое значение преобразоваывается в string
	case int:
		return strconv.Itoa(v)
	case string:
		return any.(string)
	default:
		return "???"
	}
}