package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверка тела ответа
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=wrongcity", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверка тела ответа
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверка тела ответа
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}
