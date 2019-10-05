package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	//this is a simple static server can be called with curl

	fmt.Println("Server is running in port 8000...")
	router := httprouter.New()

	// Mapping methods
	router.ServeFiles("/files/*filepath", http.Dir("/home/corganfuzz/go/src/github.com/corganfuzz/fileServer/files"))
	log.Fatal(http.ListenAndServe(":8000", router))

}
