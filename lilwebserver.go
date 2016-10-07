package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Clean(r.URL.Path)
		filePath = strings.TrimPrefix(filePath, "/")
		filePath = strings.TrimPrefix(filePath, "..")
		filePath = path.Join(workingDir, filePath)

		fmt.Printf("%s\n", filePath)

		http.ServeFile(w, r, filePath)
	})

	fmt.Printf("Serving %s on port %d...\n", workingDir, port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
