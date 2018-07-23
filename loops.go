package main
import "fmt"

func main() {
  i:= 1
  for i <= 3 {
    fmt.Println(i)
    i ++
  }

  for n:= 0; n < 20; n=n+2 {
    if n % 5 == 0 {
      continue
    } else {
      fmt.Println(n)
    }
  }
}