package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  file,_ := os.Open("sdf.txt")
  scanner := bufio.NewScanner(file)
  str := ""
  strInQue := ""
  for i:=0;i<100;i++ {
    scanner.Scan()
    strInQue = scanner.Text()
    if strInQue == "" {
      break
    } else {
      str += strInQue + string(10)
    }
  }
  fmt.Print(str)
}
