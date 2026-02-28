package main
import ("fmt")

func tampilkanCondition() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
	if time < 5 {
		fmt.Println("Good shubuh.")
	}
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}

func tampilkanSwitch() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  case 8, 9:
	fmt.Println("Ghost day")
  default:
    fmt.Println("Not a weekday")
  }
}