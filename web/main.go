package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the about home handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page ")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a About page and 2+2 is %d", sum))
}

func addValues(x, y int) (int, error) {
	s := x + y
	return s, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divide(10.0, 0)
	if err != nil {
		fmt.Fprintf(w, "can not divide by 0")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Divide of 10.0 and 10.0  is %f", f))

}

func divide(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("cannot deivide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// // main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		n, err := fmt.Fprintf(w, "Hello, world!")
// 		if err != nil {
// 			fmt.Println(err)

// 		}
// 		fmt.Println(fmt.Sprintf("Number of ytes written: %d", n))
// 	})

// 	_ = http.ListenAndServe(":8080", nil)
// }
