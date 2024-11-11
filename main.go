package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	// Obtiene el puerto desde la variable de entorno `PORT`
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Usa el puerto 8000 si `PORT` no está configurado
	}

	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New repetition")
	io.WriteString(w, "This is the fifth project in docker in the Go language")
}
