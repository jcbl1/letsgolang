package main

import (
  "fmt"
  "math"
)

func main() {
  var sum float64
  for i:=0.;i<100;i++{
    sum += (4/(8*i+1)-2/(8*i+4)-1/(8*i+5)-1/(8*i+6))/math.Pow(16,i)
    fmt.Println(sum)
  }
}
