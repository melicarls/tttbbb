package main

import (
  "log"
  "net/http"
)

func indexHandler() {
  
}

func main() {
  http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
