package main

import "fmt"

func main() {
	favorite_sport := "soccer"
	switch favorite_sport {
	case "soccer":
		fmt.Println("Go to England")
	case "surfing":
		fmt.Println("go to Hawaii")
	case "basketball":
		fmt.Println("go to whatever")
	case "swimming":
		fmt.Println("go to the pool")
	}

}
