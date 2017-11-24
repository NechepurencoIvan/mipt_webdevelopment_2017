package main

import "strings"

func RemoveEven(arr []int) []int {
  ans := []int{}
  for _, val := range(arr){
    if val % 2 == 1 {
      ans = append(ans, val)
    }
  }
  return ans
}


func PowerGenerator(e int) (func() int) {
  foundation := e
  exp := 1
  return func() (ret int) {
      exp *= foundation
      return exp
  }
}

func DifferentWordsCount(s string) int {
  s = strings.ToUpper(s)
  arr := []byte{}
  for ind, _ := range(s){
    if(strings.Compare(s[ind: ind+1], "Z") <= 0 && strings.Compare(s[ind: ind+1], "A") >= 0){
      arr = append(arr, s[ind])
    } else
    if(strings.Compare(s[ind: ind+1], "z") <= 0 && strings.Compare(s[ind: ind+1], "a") >= 0){
      arr = append(arr, s[ind])
    } else {
      arr = append(arr, '.')
    }
  }
  st := string(arr)
  p :=  strings.Split(st, ".")
  ans := len(p)
  for i := 0; i < len(p); i++ {
    if p[i] == ""{
      ans--
    } else {
      for j := i + 1; j < len(p); j++ {
        if(p[i] == p[j]){
          ans--
          break
        }
      }
    }
  }
  return ans
}
