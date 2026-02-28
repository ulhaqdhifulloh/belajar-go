package main
import ("fmt")

func myMessage(jumlahRumah int, statusRumah string) {
  fmt.Println("Aku punya rumah:", jumlahRumah, statusRumah)
}

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func tampilkanFunction() {
  myMessage(10, "nyicil")
  myMessage(20, "lunas")
  myMessage(30, "belum bayar")
  fmt.Println(myFunction(1, 2))
  testcount(1)
}