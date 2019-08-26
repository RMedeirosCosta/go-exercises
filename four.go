package main

import "fmt"

type carambolas int

var x carambolas

func main() {
	fmt.Println(x)
	fmt.Println("%T", x)
	x = 42
	fmt.Println(x)
}
