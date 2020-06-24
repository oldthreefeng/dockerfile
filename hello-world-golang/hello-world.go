package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.Method, r.URL, r.Host, r.RemoteAddr)
		version := os.Getenv("VERSION")
		if version == "" {
			version = "v1.0.1"
			// version = "v2" 模拟发布. v3 v4 v5 v6
		}
		fmt.Fprintf(w, "Hello Kubernetes ! Hello World version: %s\n", version)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
