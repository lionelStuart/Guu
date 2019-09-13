package main

import (
	. "Guu/p3/base1"
	"net/http"
)

func main() {

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
