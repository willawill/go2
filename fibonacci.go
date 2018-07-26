package main

import "fmt"

func fibonacci() func() int{
  a, b := 0, 1
  return func() (c int) {
    c = a
    a, b = b, a + b
    return
  }
}

func main() {
  f:= fibonacci()
  for i:= 0; i< 5; i++ {
    fmt.Println(f())
  }
}