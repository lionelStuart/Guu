package base1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	engine := New()
	engine.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>hello guu</h1>")
	})

	engine.GET("/hello", func(c *Context) {
		c.String(http.StatusOK, "hello %s, you're at %s \n", c.Query("name"), c.Path)
	})

	engine.POST("/login", func(c *Context) {
		c.Json(http.StatusOK, H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	engine.Run(":9999")
}

func TestClient_login(t *testing.T) {
	resp, err := http.PostForm("http://localhost:9999/login",
		url.Values{"username": {"jim"}, "password": {"1234"}})
	defer resp.Body.Close()
	if err != nil {
		t.Fatal("fail", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("result body:%s\n", string(body))

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal("failure parse body ", body)
	}
	fmt.Printf("result user:%+v\n", user)

}
