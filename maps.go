package main
import ("fmt")

func tampilkanMap() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

func tampilkanMap2() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"
  a["condition"] = "broken"

  fmt.Println(a)

  a["year"] = "1970" // Updating an element
  a["color"] = "red" // Adding an element
  delete(a, "condition")

  fmt.Println(a)
}

func tampilkanMap3() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var b []string             // defining the order
  b = append(b, "one", "two", "three", "four")

  for k, v := range a {        // loop with no order
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }
}