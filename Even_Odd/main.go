package main

import "fmt"

func main() {

	newList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(newList)
	for _, i := range newList {
		if i%2 == 0 {
			fmt.Println("Even", i)
		} else {
			fmt.Println("Odd", i)
		}
	}

}
