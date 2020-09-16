package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	url := req.URL
	log.Print("---", "Requester: ", req.RemoteAddr, "  ---", "Method: ", req.Method, "  ---", "Requested URI: ", url.Path)

	fmt.Fprintln(resp, `<body style="background-color:powderblue;">`)
	fmt.Fprintln(resp, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalf("Error listening: ", err)
	}
}
