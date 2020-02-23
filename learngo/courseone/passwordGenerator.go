package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
  "os"
  "bufio"
)

func toNum(s []byte) int {
  var res float64
  for i,v := range s {
    if i == len(s) - 2 {
      break
    }
    shrt := len(s) - i - 3
    res += (float64(v) - 48) * math.Pow(10, float64(shrt))
  }
  return int(res)
}

func main() {
  var o1, o2 byte
  var o3 []byte
  var err error
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Println("是否包含数字？（Y/n）")
    for {
      o1,err = reader.ReadByte()
      if err != nil {
        continue
      } else {
        if o1 == 89 || o1 == 121 {
          o1 = 1
          break
        } else if o1 == 78 || o1 == 110 {
          o1 = 0
          break
        } else {
          fmt.Println("请输入正确的选项（Y/n）")
        }
      }
    }
    reader.Reset(os.Stdin)
    fmt.Println("是否包含字母？（Y/n）")
    for {
      o2,err = reader.ReadByte()
      if err != nil {
        continue
      } else {
        if o2 == 89 || o2 == 121 {
          o2 = 1
          break
        } else if o2 == 78 || o2 == 110 {
          o2 = 0
          break
        } else {
          fmt.Println("请输入正确的选项（Y/n）")
        }
      }
    }
    reader.Reset(os.Stdin)
    if o1 == 0 && o2 ==0 {
      fmt.Println("请至少允许一项")
      continue
    } else {
      break
    }
  }
  fmt.Print("请输入密码的长度：")
  o3,_ = reader.ReadBytes(10)
  reader.Reset(os.Stdin)
  rand.Seed(time.Now().UnixNano())
  passwd := ""
  for i:=0;;i++{
    n := rand.Intn(94) + 33
    if len(passwd) >= toNum(o3) {
      break
    }
    if n >= 48 && n <= 57 && o1 == 1 {
      passwd += string(n)
    } else if n >= 65 && n <= 90 && o2 == 1 {
      passwd += string(n)
    } else if n >= 97 && n <= 122 && o2 == 1 {
      passwd += string(n)
    }
  }
  fmt.Println(passwd)
  fmt.Print("请保存这个密码后按回车键退出此程序")
  reader.ReadString(10)
}
