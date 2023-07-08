package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	router := NewRouter()
	fmt.Println("The server is running on port 8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)	
}

