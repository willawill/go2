package main

import "fmt"

func mergeSort(A []int) []int {
  if len(A) <= 1 {
    return A
  }
  mid := len(A)/2
  return merge(mergeSort(A[:mid]), mergeSort(A[mid:]))
}

func merge(left, right []int) []int {
  ret := make([]int, (len(left)+len(right)))

  for len(left) > 0 || len(right) > 0 {
    if len(left) == 0 {
      return append(ret, right...)
    }
    if len(right) == 0 {
      return append(ret, left...)
    }
    if left[0] <= right[0] {
      ret = append(ret, left[0])
      left = left[1:]
    } else {
      ret = append(ret, right[0])
      right = right[1:]
    }
  }
  return ret
}

func main() {
  s:= []int{3, 41, 52, 26, 38, 57, 9, 49}
  fmt.Println(mergeSort(s))
}