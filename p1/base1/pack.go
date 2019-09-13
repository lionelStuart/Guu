package base1

import (
	"fmt"
	"net/http"
)

func TestBase1() {
	fmt.Println("test base 1")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %s \n", k, v)
	}
}

//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func ListenAndServe(address string, h Handler) error {
	return nil
}

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	case "/index":
		fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s \n", r.URL)
	}
}
