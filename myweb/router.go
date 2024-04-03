package myweb

import "net/http"

// router结构体
type router struct {
	handlers map[string]HandlerFunc
}

// 创建实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 添加路由
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	r.handlers[method+"-"+pattern] = handler
}

// 路由处理方法
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
