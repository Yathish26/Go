package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World this is Made using Go Server")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting Server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error Starting Server: ", err)
	}

}
