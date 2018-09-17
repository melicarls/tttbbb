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

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", homepageHandler).Methods("GET")
  return r
}

func main() {
  router := newRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
