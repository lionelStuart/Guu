package base1

import (
	"net/http"
	"testing"
)

func TestS1(t *testing.T) {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	t.Fatal(http.ListenAndServe(":9999", nil))
}

func TestS2(t *testing.T) {
	engine := &Engine{}
	t.Fatal(http.ListenAndServe(":9999", engine))
}
