package main 
import "fmt"

func main() {
    i := 1
// While Syntax
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
// For loop syntax
    for i :=1; i <= 10; i++ {
        fmt.Println(i)
    
    }
// if variant
for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
        fmt.Println(i, "even")
    } else {
        fmt.Println(i, "odd")
    }
}
// switch case
i = 1
for i<=10 {
switch i {
case 0: fmt.Println("zero")
case 1: fmt.Println("One")
case 2: fmt.Println("Two")
case 3: fmt.Println("Three")
case 4: fmt.Println("Four")
case 5: fmt.Println("Five")
default: fmt.Println("Unknown number")
    }
    i += 1
}

}

