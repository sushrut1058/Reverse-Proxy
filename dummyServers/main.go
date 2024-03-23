package main

import (
	"fmt"
	"net/http"
	"os"
)

var response string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[PORT: %s] Request received on server\n", os.Args[1])
	resp := fmt.Sprintf("Hello from server at Port:%v\n, Also server says: %v\n", os.Args[1], response)
	fmt.Fprintf(w, resp)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No port passed in args")
		return
	}
	file, err := os.ReadFile("body.txt")
	if err != nil {
		fmt.Println("[Server] Error reading file", err)
	}

	response = string(file)

	http.HandleFunc("/", handler)
	port := os.Args[1]
	fmt.Printf("Starting on port: %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error received in ListenAndServe: %s\n", err)
	}
}
