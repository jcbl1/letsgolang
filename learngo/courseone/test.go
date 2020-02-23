package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  b,_ := reader.ReadByte()
  fmt.Printf("%T %v", b, b)
}
