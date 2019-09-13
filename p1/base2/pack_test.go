package base2

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEngine_ServeHTTP(t *testing.T) {
	engine := New()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	})

	engine.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	})

	t.Fatal(engine.Run(":9999"))

}
