package mid

import (
	. "Guu/p5/base1"
	"net/http"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	engine := New()
	engine.Use(Logger())
	engine.GET("/index", func(ctx *Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := engine.Group("v1")
	v1.GET("/", func(ctx *Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v1.GET("/hello", func(ctx *Context) {
		ctx.String(http.StatusOK, "hello %s ,you're at %s", ctx.Query("name"),
			ctx.Path)
	})

	v2 := engine.Group("v2")
	v2.Use(MidForV2())
	v2.GET("/hello/:name", func(ctx *Context) {
		ctx.String(http.StatusOK, "hello %s ,you're at %s", ctx.Query("name"),
			ctx.Path)
	})
	v2.POST("/login", func(ctx *Context) {
		ctx.Json(http.StatusOK, H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	engine.Run(":9999")
}
