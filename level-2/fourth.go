package main

import "fmt"

func main() {
	x := 13
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	shifted_x := x << 1
	fmt.Printf("%d\n%b\n%#x\n", shifted_x, shifted_x, shifted_x)
}
