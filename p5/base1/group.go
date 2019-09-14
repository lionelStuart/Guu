package base1

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
