package myweb

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc是处理程序
type HandlerFunc func(c *Context)

// RouterGroup和Engine定义
type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		//Group对象要有访问Router能力，定义指向它的指针，从而所有资源统一由其协调和访问
		engine *Engine
	}
	Engine struct {
		//嵌入，可访问RouterGroup中字段而无需命名
		*RouterGroup
		router *router
		//Engine作为最顶层分组，拥有RouterGroup所有能力，存储所有groups
		groups []*RouterGroup
	}
)

// 创建实例
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 创建分组，由Engine统一管理，因此所有分组都共享统一Engine实例
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	routerGroup := &RouterGroup{
		prefix: g.prefix + prefix, //注意这里需加上g.prefix
		parent: g,                 //画图就明白，把两个结构以图的形式画出
		engine: g.engine,          //共享统一Engine实例
	}
	//这里要append到g.engine.groups，因为所有分组都共享统一Engine实例
	g.engine.groups = append(g.engine.groups, routerGroup)
	return routerGroup
}

// 添加路由，这里没使用parent，之前设计时后面简化，所以RouterGroup的parent可去掉
func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRouter(method, pattern, handler)
}

// 添加GET路由，注意改成g
func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

// 添加POST路由，注意改成g
func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// 添加中间件到group中
func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// 启动web服务
func (engine *Engine) Run(addr string) (err error) {
	// go中实现某个接口方法的类型都可自动转换为某个接口类型
	//下面等价于return http.ListenAndServe(addr,(http.Handler)(engine))
	return http.ListenAndServe(addr, engine)
}

// 只要传入任何实现ServerHTTP接口的实例，所有HTTP请求，都由该实例处理，注意中间件要在请求前后。
// 中间件要应用到某个Group
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
