package main
import (
    "fmt";
    "math"
)
func square(x float64) float64 {
    return x*x
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(square(a) + square(b))
}

/**

func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}

func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r
}
**/
// Create Structs and stuffs for all this
type Circle struct {
    x, y, r float64 
}
func (c *Circle) area() float64 {
    return math.Pi * square(c.r)
}
type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func main() {
    c := Circle{x:0, y: 0, r: 5}
    // or c := Circle{0, 0, 5} if you remember order of variables
    fmt.Println("Circle:")
    fmt.Println(c.x, c.y, c.r)
    fmt.Println(c.area()) // this happens because of area() method, otherwise cicleArea(&c)
}
