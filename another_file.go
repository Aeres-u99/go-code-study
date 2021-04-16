package main

import (
	"fmt"
)

func main() {
	runeArray := []rune{'a', 'b', '$'}
	s := string(runeArray)
	fmt.Println(s)
	t := "ab$"
	r := []rune(t)
	fmt.Printf("%U\n", r)
}
