package main
import "fmt"

func binarySearch(arr []int, item int) (int, bool) {
  l := len(arr)
  low := 0
  high := l - 1

  for (low < high) {
    fmt.Println(low, high)
    mid := low + (high - low)/2

    if arr[mid] == item {
      return mid, true
    } else if arr[mid] > item {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  return 0, false
}

func bsR(arr []int, low int, high int, item int) (int, bool) {
  mid := low + (high - low)/2
  
  for (low <= high) {
    if arr[mid] ==  item {
      return mid, true
    }
    if arr[mid] > item {
      return bsR(arr, low, mid-1, item)
    } 
    return bsR(arr, mid+1, high, item)
  }  
  return 0, false

}
func main() {
  s := []int{1, 2, 3, 4, 5, 6}

  fmt.Println(binarySearch(s, 5))
  fmt.Println(bsR(s, 0, 6, 5))
}