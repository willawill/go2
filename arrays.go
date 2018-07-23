package main
import "fmt"

// slices and arrays
// slices is like dynamic assigned array but the type is decided by the element it contained
// arrays are pass by value whereas slices are pass by reference
func main() {
  var a [5]int
  b := make([]string, 5)

  a[1] = 5
  fmt.Println(a, a[1])
  b[0] = "Apple"
  b[1] = "Pearl"
  b = append(b, "Cherry", "Potato")
  fmt.Println(b)

  // make a copy
  c := make([]string, len(b))
  copy(c, b)
  fmt.Println(c)
}

