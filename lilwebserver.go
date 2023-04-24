package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	workingDir := "."
	if len(os.Args) > 1 {
		workingDir = os.Args[1]
	}

	port := 8080
	if len(os.Args) > 2 {
		var err error
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid port number '%s'\n", os.Args[2])
			os.Exit(1)
		}
	}

	handler := http.FileServer(http.Dir(workingDir))

	fmt.Printf("Serving %s on port %d...\n", workingDir, port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), handler)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
