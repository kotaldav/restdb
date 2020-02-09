package main

import (
  "fmt"
  "net/http"
  
  "github.com/gorrila/mux"
)

var (
  
)

func dbInit() {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world!")
}

func main() {

  router := mux.NewRouter()
  router.HandleFunc("/", homeHandler)

  http.Handle("/", router)
}
