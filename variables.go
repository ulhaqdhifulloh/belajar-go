// Belajar variabel, nilai konstan, output

package main

import ("fmt")

var b = 10
const hours = 24 // unchangable 
const (
    nama = "Ulhaq"
    hobi string = "Scroll Fesnuk"
)

func tampilkanVariabel() {
    var student1 string = "John"
    student2 := "Jane"
    studentId := 2
    var a bool
    student2 = "Janetes"
    var x, y, z int = 10, 15, 20
    var c, d = 10, "huhu"
    var A = true
    
    fmt.Println("Nama Siswa 1:", student1)
    fmt.Println("Nama Siswa 2:", student2)
    fmt.Println("ID Kelas:", studentId)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(x+y+z+c)
    fmt.Println(d)
    fmt.Println(A)
    fmt.Printf("Dalam satu hari terdapat %v jam", hours)
    fmt.Printf("Tipe data %v adalah %T", hobi, hobi)
    fmt.Printf("%t kah?", A)
    fmt.Print(nama, hobi)
    fmt.Println(nama, hobi)
}