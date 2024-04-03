package myweb

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc是处理程序
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// Engine定义
type Engine struct {
	router map[string]HandlerFunc
}

// 创建实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[method+"-"+pattern] = handler
}

// 添加GET路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 添加POST路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 启动web服务
func (engine *Engine) Run(addr string) (err error) {
	// go中实现某个接口方法的类型都可自动转换为某个接口类型
	//下面等价于return http.ListenAndServe(addr,(http.Handler)(engine))
	return http.ListenAndServe(addr, engine)
}

// 只要传入任何实现ServerHTTP接口的实例，所有HTTP请求，都由该实例处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}
