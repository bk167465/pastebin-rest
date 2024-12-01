package main

import (
   "fmt"
   "net/http"
   routes "github.com/bk167465/paste-bin/routes"
)

func main() {
   http.HandleFunc("/paste", routes.CreateNewPaste)
   http.HandleFunc("/getPaste", routes.GetPaste)
   fmt.Println("server is starting at 8080")
   http.ListenAndServe(":8080", nil)
}