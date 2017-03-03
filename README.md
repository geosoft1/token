token
====

Package token provide a random number generator

     fmt.Println(token.GetToken(6))

Web service example

     package main
     
     import (
     	"fmt"
     	"net/http"
     
     	"github.com/geosoft1/token"
     )
     
     func main() {
     	http.HandleFunc("/",
     		func(w http.ResponseWriter, r *http.Request) {
     			fmt.Fprintf(w, token.GetToken(6))
     		})
     	http.ListenAndServe(":8080", nil)
     }
