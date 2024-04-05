package myweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 方便构建JSON数据
type H map[string]interface{}

// Context结构体，封装请求和响应、请求的Method、Path和响应的响应码、中间件
type Context struct {
	//objects
	Request *http.Request
	Writer  http.ResponseWriter
	//request info
	Path   string
	Method string
	Params map[string]string //新增对象，存储解析到的参数
	//response info
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int
}

// 获取存储的解析到的参数，如/hello/:name,/hello/attackoncs
// 解析存储到的c.Params["name"]="attackoncs"
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 创建Context实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Request: req,
		Writer:  w,
		Path:    req.URL.Path,
		Method:  req.Method,
		index:   -1,
	}
}

// 中间件调下个流程，注意这里要遍历所有handler，原因是不是所有的handler都会调Next，手动调Next
// 一般用于在请求前后各实现一些行为。若只作用请求前，则可省略掉Next,算是兼容性比较好的写法。
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// 填充请求中的key
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// 查询请求中的url的key
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 设置响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 构造String响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 构造JSON响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 构造Data响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 构造HTML响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
