package main

import "fmt"

const (
	orig = iota
	asm
	golang
)

func main() {
	// Initialization
	var data *[25]uint64
	var _data [25]uint64
	data = &_data
	init_data(data)
	pretty_print(data, orig)

	// Run the operation in assembly
	consecutively_add(data)
	pretty_print(data, asm)

	// Run the operation in Golang
	init_data(data)
	for i := 0; i < 24; i += 2 {
		data[i] = data[i] + data[i+1]
	}
	pretty_print(data, golang)
}

func init_data(data *[25]uint64) {
	for i := 0; i < 25; i++ {
		var j uint64 = uint64(i)
		data[i] = j
	}
}

func pretty_print(data *[25]uint64, kind int) {
	switch kind {
	case orig:
		fmt.Print("orig:   ")
		for i := 0; i < 25; i++ {
			if i == 24 {
				fmt.Print(data[i], "\n")
			} else if i == 6 || i == 9 {
				fmt.Print(" ", data[i], ", ")
			} else {
				fmt.Print(data[i], ", ")
			}
		}
	case asm:
		fmt.Print("asm:    ")
		for i := 0; i < 25; i++ {
			if i == 24 {
				fmt.Print(data[i], "\n")
			} else {
				fmt.Print(data[i], ", ")
			}
		}
	case golang:
		fmt.Print("golang: ")
		for i := 0; i < 25; i++ {
			if i == 24 {
				fmt.Print(data[i], "\n")
			} else {
				fmt.Print(data[i], ", ")
			}
		}
	}
}
