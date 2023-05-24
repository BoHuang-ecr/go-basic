/*
	The Go language provides another data type, the interface, which defines all common methods together.
	As long as any other type implements these methods, it implements this interface.
	Interfaces allow us to bind different types to a common set of methods, enabling polymorphism and flexible design.
	Interfaces in Go are implicitly implemented, that is, if a type implements all the methods defined by an interface, then it automatically implements that interface.
	Therefore, we can achieve polymorphism by passing interfaces as parameters to implement calls to different types.
*/

package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle Area: %f\n", s.area())

	s = Circle{radius: 3}
	fmt.Printf("Circle Area: %f\n", s.area())
}
