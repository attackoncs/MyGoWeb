package main

import (
	"myweb"
	"net/http"
)

func main() {
	engine := myweb.Default()
	engine.GET("/", func(c *myweb.Context) {
		c.String(http.StatusOK, "hello attackoncs")
	})
	//边界越界
	engine.GET("/panic", func(c *myweb.Context) {
		names := []string{"attackoncs"}
		c.String(http.StatusOK, names[100])
	})

	engine.Run(":9999")
}
