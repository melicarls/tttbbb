package main

import (
  "html/template"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
  templates := template.Must(template.ParseFiles("views/layout.html", "views/birdview.html"))
  templates.ExecuteTemplate(w, "layout", nil)
}

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", homepageHandler).Methods("GET")
  r.HandleFunc("/birds", getBirdHandler).Methods("GET")
  r.HandleFunc("/birds", createBirdHandler).Methods("POST")

  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

  r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

  return r
}

func main() {
  router := newRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
