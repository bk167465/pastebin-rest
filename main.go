package main

import (
   "fmt"
   "net/http"
   routes "github.com/bk167465/paste-bin/routes"
)

func main() {
   http.HandleFunc("/paste", routes.CreateNewPaste)
   fmt.Println("Server is running on PORT :8080")
   http.ListenAndServe(":8080", nil)
}