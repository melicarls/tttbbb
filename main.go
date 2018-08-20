package main

import (
  "html/template"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("views/homepage.html"))

func homepageHandler(w http.ResponseWriter, r *http.Request) {
  templates.ExecuteTemplate(w, "homepage.html", nil)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", homepageHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}
