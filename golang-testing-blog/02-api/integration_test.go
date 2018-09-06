package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeTodoRequest struct{}

func TestGetProudTodoTitle(t *testing.T) {
	expected := "SWEET!!!"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"title\": \"sweet\"}")
	}))
	resp := getProudTodoTitle(ts.URL)
	if resp != expected {
		t.Errorf("Fail")
	}
}
