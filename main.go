package main

import "fmt"

var x string = "Super helper string, Accessible to everyone"

func main() {
	i := 1
	chars := 0
	var lines int 
	fmt.Scanf("%d", &lines)
    fmt.Println("---")
    fmt.Println(lines)
    for i <= lines {
        chars = 0
		for chars <= i {
			fmt.Print("* ")
            chars += 1
		}
		fmt.Println("")
        i += 1
	}
    
    for i :=1; i <= lines ; i++ {
        fmt.Println("Line number: ", i)
    }
}

func f() {
	fmt.Println(x)
}
