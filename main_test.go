// main_test.go
package main

import (
	"apiusuarios/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsuariosHandler(t *testing.T) {
	// Crear una solicitud HTTP falsa
	req, err := http.NewRequest("GET", "/usuarios", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un ResponseRecorder para registrar la respuesta
	rr := httptest.NewRecorder()

	// Llamar al manejador con la solicitud falsa
	handlers.GetUsuariosHandler(rr, req)

	// Verificar el código de estado de la respuesta (debería ser 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("código de estado inesperado: obtuve %v, quiero %v", rr.Code, http.StatusOK)
	}

	// Verificar el contenido de la respuesta (puedes implementar esto según tu formato JSON esperado)
	// Por ejemplo, puedes verificar si la respuesta contiene ciertos campos o valores específicos.
	// Aquí, solo estamos comprobando si la respuesta no está vacía.
	if rr.Body.String() == "" {
		t.Error("La respuesta está vacía")
	}
}
