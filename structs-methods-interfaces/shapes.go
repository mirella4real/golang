package main

import "fmt" 
import "math"


// A struct is just a named collection of fields where you can store data.
type Rectangle struct {
    Width float64
    Height float64
}

//A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

//​Interfaces allow you to make functions that can be used with different types and create highly-decoupled code 
//In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile
type Shape interface {
    Area() float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width  + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}

func main() {
    var rectangle Rectangle
    rectangle.Width = 300
    rectangle.Height = 200
    fmt.Printf("the rectangle's dimensions are %2.f  by %2.f  \n", rectangle.Width, rectangle.Height)
}