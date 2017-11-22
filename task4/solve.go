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
