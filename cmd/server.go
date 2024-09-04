package main

import (
	"fmt"
	"net/http"
	"os"

	flymetarutility "github.com/daniloraisi/fly-metar-utility/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", flymetarutility.Handler)

	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
