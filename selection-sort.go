package main
import "fmt"

func selectionSort(A []int) []int {
  for j:= 0; j < len(A); j++ {
    min := j
    for i := j + 1; i < len(A); i++ {
      if A[i] >= A[min] {
        continue
      }
      min = i
    }
    if min != j {
      A[min], A[j] = A[j], A[min]
    }
  }
  return A
}

func main() {
  s := []int{1, 3, 2, 4, 8,6}

  fmt.Println(selectionSort(s))
} 