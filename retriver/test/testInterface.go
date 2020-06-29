package main
import (
	"fmt"
)
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Skin interface {
	Color() float64
}

type Cube struct {
	side float64
}

func (c Cube)Area() float64 {
	return c.side * c.side
}

func (c Cube)Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	// type assertion
	var s Shape = Cube{3.0}
	value1, ok1 := s.(Object)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Object? %v\n", value1, ok1)
	value2, ok2 := s.(Skin)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Skin? %v\n", value2, ok2)

	var obj Object = Cube{2}
	value3, ok3 := obj.(Shape)
	fmt.Printf("dynamic value of Object 'obj' with value %v implements interface Shape? %v\n", value3, ok3)
}

