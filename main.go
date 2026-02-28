package main

import (
	"fmt"
	"strings"
)

var gate = "\n" + strings.Repeat("*", 20) + "\n"

func main() {
	fmt.Println("--- Aplikasi Belajar Go Dimulai ---")
	fmt.Print(gate)
	tampilkanHelloWorld()
	fmt.Print(gate)
	tampilkanVariabel()
	fmt.Print(gate)
	tampilkanTipeData()
	fmt.Print(gate)

	fmt.Println("--- Program Selesai ---")
}
