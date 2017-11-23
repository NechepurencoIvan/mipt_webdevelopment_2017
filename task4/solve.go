package main

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
