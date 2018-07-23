package main
import "time"
import "fmt"

func main() {
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("Have Fun!")
  case time.Monday: 
    fmt.Println("Go back to work")
  default: 
    fmt.Println("Whatever today is")
  }

  whereami := func(i interface{}) {
    switch i.(type) {
      case int:
        fmt.Println("It is an int")
      case string:
        fmt.Println("It is a string")
      default:
        fmt.Println("I don't know")
    }
  }

  whereami(1)
}