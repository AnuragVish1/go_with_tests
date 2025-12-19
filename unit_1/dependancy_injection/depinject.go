package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":4321", http.HandlerFunc(GreetHandler)))
}

func Greet(writer io.Writer, word string) {
	fmt.Fprintf(writer, "Hello, %s", word)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Anruag")
}
