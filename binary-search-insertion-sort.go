package main

import "fmt"

func insertionSort(input []int) []int {
  for i := 1; i < len(input); i++ {
    j := i
    for j > 0 && input[j - 1] > input[j] {
      input[j - 1], input[j] = input[j], input[j - 1]
      j = j - 1
    }
  }
  return input
}

func insertEIntoSortedArray(input []int, item int) []int {
// find the spot to insert the element and move everything backward
  low := 0
  high := len(input) - 1
  result := make([]int, 1)
  if item < input[0] {
    result := make([]int, len(input))
    copy(result, input)
    result = append([]int{item}, result...)
    return result
  }

  if item > input[high] {
    result := make([]int, len(input))
    copy(result, input)
    result = append(result, item)
    return result
  }

  for low < high {
    mid := low + (high - low)/2
    if input[mid] < item && input[mid + 1] > item{
      result = append(input[0:mid+1], item)
      result = append(result, input[mid+1:]...)
      return result
    }

    if input[mid] < item {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }
  return result
}

func main() {
  s:= []int{26, 6, 31, 10, 15, 6}
  fmt.Println(insertionSort(s))
  fmt.Println(insertEIntoSortedArray(insertionSort(s), 5))
  fmt.Println(insertEIntoSortedArray(insertionSort(s), 11))
  fmt.Println(insertEIntoSortedArray(insertionSort(s), 35))
}