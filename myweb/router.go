package myweb

import (
	"net/http"
	"strings"
)

// router结构体
// roots key 如：roots['GET'] roots['POST']
// handlers key 如：handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']
type router struct {
	roots    map[string]*node       //存储每种请求的Trie树根节点
	handlers map[string]HandlerFunc //每种请求方式的HandlerFunc
}

// 创建实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

// 解析路由，如/hello/b/c，注意只有一个*被允许
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			//匹配到第一个*就跳出
			if item == "*" {
				break
			}
		}
	}
	return parts
}

// 添加路由，这里会先用/分割pattern，并插入到对应方法（GET、POST）trie树中，设置handler
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 返回路由，返回路由解析到的服务节点和存储去除":"和"*"后的部分分段对应的path的哈希表
// 如/p/go/doc匹配到/p/:lang/doc，解析的哈希表结果为：{lang: "go"}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// 路由处理方法，获取到路由节点和实际解析到的参数
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
