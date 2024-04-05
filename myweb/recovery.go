package myweb

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 获取堆栈，调runtime.Callers返回调用栈程序计数器
func trace(message string) string {
	var pcs [32]uintptr
	//跳过前3Caller（第0Caller是Callers本身，第1是上层trace，第2是再上一层的defer func）
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)   //获取对应的函数
		file, line := fn.FileLine(pc) //取到调用该函数的文件名和行号
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// Recovery()中间件，defer挂载错误恢复函数，调recover捕获panic，将堆栈信息打印在日志中
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
