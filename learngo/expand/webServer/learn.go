package main

import (
  "fmt"
  "net/http"
  "os"
  "bufio"
)

var fileStr, markdownIndex, strInQue string

func HomeFunc(rw http.ResponseWriter, rq *http.Request) {
  fmt.Fprint(rw, fileStr)
}
func MarkdownFunc(rw http.ResponseWriter, rq *http.Request) {
  fmt.Fprint(rw, markdownIndex)
}

func main() {
  homepage,err := os.Open("index.html")
  defer homepage.Close()
  if err != nil {
    fmt.Println(err)
    for {}
  }
  scanner := bufio.NewScanner(homepage)
  for i:=0;; {
    scanner.Scan()
    strInQue = scanner.Text()
    if strInQue == "" {
      if i >=5 {
        break
      } else {
        i++
        fileStr += strInQue + string(10)
      }
    } else {
      i = 0
      fileStr += strInQue + string(10)
    }
  }
  markdownfile,err := os.Open("index")
  defer markdownfile.Close()
  if err != nil {
    fmt.Println(err)
    for {}
  }
  scanner = bufio.NewScanner(markdownfile)
  for i:=0;; {
    scanner.Scan()
    strInQue = scanner.Text()
    if strInQue == "" {
      if i >=5 {
        break
      } else {
        i++
        markdownIndex += strInQue + string(10)
      }
    } else {
      i = 0
      markdownIndex += strInQue + string(10)
    }
  }
  http.HandleFunc("/", HomeFunc)
  http.HandleFunc("/markdown/", MarkdownFunc)
  fmt.Println("Starting")
  http.ListenAndServe(":23337", nil)
  fmt.Println("Server started.")
}
