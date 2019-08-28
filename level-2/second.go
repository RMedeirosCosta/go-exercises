package main

import "fmt"

const carambolas = "carambolas"

func main() {
	equal := (1 == 1)
	less_or_equal := (1 <= 2)
	greater_or_equal := (2 >= 1)
	different := (carambolas != "batatinhas")
	same := (carambolas == carambolas)
	less := (1 < 2)
	greater := (2 > 1)

	fmt.Println(equal, less_or_equal, greater_or_equal, different, same, less, greater)
}
