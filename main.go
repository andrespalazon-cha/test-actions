package main

import (
	"fmt"
	"net/http"
)

// helloHandler responde a las peticiones HTTP
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola! La API y el CI/CD funcionan correctamente.")
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Servidor iniciando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
