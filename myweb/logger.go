package myweb

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		//开始时间
		t := time.Now()
		//处理请求
		c.Next()
		//计算所用时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
