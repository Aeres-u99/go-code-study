package main
import "fmt"

func swap(x *int,y *int) {
   var temp int
   temp = *y
   *y = *x
   *x = temp
}

func main() {
    x := 5
    y := 6
    fmt.Println(x,y)
    swap(&x,&y)
    fmt.Println(x,y)
}
