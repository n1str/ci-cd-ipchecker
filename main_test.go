package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fakeIPAPIHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "123.45.67.89")
}

func TestGetIP(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(fakeIPAPIHandler))
	defer testServer.Close()

	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	expectedIP := "123.45.67.89"
	if string(body) != expectedIP {
		t.Errorf("Ожидался IP %s, но был получен %s", expectedIP, string(body))
	}
}
