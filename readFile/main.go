package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct {
}

// func main() {
// 	fileName := os.Args[1]
// 	fmt.Println(fileName)
// 	File, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// 	lw := logWriter{}
// 	io.Copy(lw, File)
// }

// func (logWriter) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	fmt.Println("Just wrote this many bytes :", len(bs))
// 	return len(bs), nil
// }

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	File, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, File)
}
