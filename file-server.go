package main

import (
  "fmt"
  "log"
  "net/http"
  flag "github.com/ogier/pflag"
)

var (
  path string
)

func main() {
  flag.Parse()
  fmt.Printf("Directory path: %s\n", path)
  // Simple static webserver:
  fmt.Printf("File server is runing at http://localhost:8080\n")
  log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(path))))
}

func init() {
  flag.StringVarP(&path, "path", "p", "", "Directory path")
}
