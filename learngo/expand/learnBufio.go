package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  fmt.Println("Hey there. What is your name?")
  reader := bufio.NewReader(os.Stdin)
  line,err := reader.ReadString(10)
  if err != nil {
    fmt.Println("!")
  } else {
    fmt.Print("Hi, ", line)
  }
}
