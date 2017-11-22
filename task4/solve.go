package main

import "fmt"

func RemoveEven(arr []int) []int {
  ans := []int{}
  for _, val := range(arr){
    if val % 2 == 1 {
      ans = append(ans, val)
    }
  }
  return ans
}


func main() {
  input := []int{0, 3, 2, 5}
  result := RemoveEven(input)
  fmt.Println(result) // Должно напечататься [3 5]
}
