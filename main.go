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
	tampilkanSlice()
	tampilkanSlice2()
	fmt.Print(gate)
	tampilkanCondition()
	fmt.Print(gate)
	tampilkanPerulangan()
	tampilkanPerulangan2()
	fmt.Print(gate)
	tampilkanFunction()
	fmt.Print(gate)
	tampilkanStruct()
	fmt.Print(gate)
	tampilkanMap()
	tampilkanMap2()
	tampilkanMap3()
	fmt.Print(gate)

	fmt.Println("--- Program Selesai ---")
}
