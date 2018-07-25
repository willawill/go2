package main
import "fmt"

func selectionSort(A []int) []int {
  B := make([]int, len(A))
  count := 0
  for count < len(A) {
    min := A[0]
    index := 0
    for i := 1; i < len(A); i++ {
      if A[i] >= min {
        continue
      }
      min = A[i]
      index = i
    }
    B[count] = min
    count++
    A[index] = 100
  }
  return B
}

func main() {
  s := []int{1, 3, 2, 4, 8,6}

  fmt.Println(selectionSort(s))
} 