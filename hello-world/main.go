package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, `<body style="background-color:powderblue;">`)
	fmt.Fprintln(resp, "Hello World! By Erik Nobel")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error listening: ", err)
	}
}
