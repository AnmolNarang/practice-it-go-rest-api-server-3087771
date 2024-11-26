package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Friends and family!!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server Started on local host 9003")
	log.Fatal(http.ListenAndServe(":9003", nil))

}
