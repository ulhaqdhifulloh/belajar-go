package main
import ("fmt")

func tampilkanArray() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}
  arr3 := [5]int{1:10,2:40}
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  var zero = [3]int{}
  cars[1] = "Ferrari"

  fmt.Println(arr1)
  fmt.Println(len(arr1))
  fmt.Println(arr2)
  fmt.Println(arr3)
  fmt.Println(zero)
  fmt.Print(cars)
  fmt.Println(cars[0] + cars[1])
}