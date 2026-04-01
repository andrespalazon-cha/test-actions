package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// 1. Creamos una petición simulada
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Creamos un grabador para capturar la respuesta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// 3. Ejecutamos el handler
	handler.ServeHTTP(rr, req)

	// 4. Verificamos el código de estado (Aceptación)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: obtuvimos %v, esperábamos %v", status, http.StatusOK)
	}

	// 5. Verificamos el cuerpo de la respuesta
	expected := "¡Hola! La API y el CI/CD funcionan correctamente."
	if rr.Body.String() != expected {
		t.Errorf("Cuerpo inesperado: obtuvimos %v, esperábamos %v", rr.Body.String(), expected)
	}
}
