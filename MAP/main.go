package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#ff0000",
		"red1": "#ff0000",
	}
	fmt.Printf("%v", colors)
	fmt.Println(colors)
}
