package base1

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (rp *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rp.engine
	newGroup := &RouterGroup{
		prefix: rp.prefix + prefix,
		parent: rp,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (rp *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(rp.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (rp *RouterGroup) Static(relativePath string, root string) {
	handler := rp.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	rp.GET(urlPattern, handler)
}

func (rp *RouterGroup) Use(middlewares ...HandlerFunc) {
	rp.middlewares = append(rp.middlewares, middlewares...)
}

func (rp *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := rp.prefix + comp
	rp.engine.router.addRoute(method, pattern, handler)
}

func (rp *RouterGroup) GET(pattern string, handler HandlerFunc) {
	rp.addRoute("GET", pattern, handler)
}

func (rp *RouterGroup) POST(pattern string, handler HandlerFunc) {
	rp.addRoute("POST", pattern, handler)
}
