package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("./public")
	handler := http.FileServer(dir)
	http.Handle("/", handler)

	log.Println("Sirviendo archivos en Http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
