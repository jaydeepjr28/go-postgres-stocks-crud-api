package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Printf("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
