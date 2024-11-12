package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola soy Esteban y este es el deber es de programacion distribuida con el lenguaje go")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8086", nil)
}