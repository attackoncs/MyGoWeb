package main

import (
	"myweb"
	"net/http"
)

func main() {
	engine := myweb.New()
	engine.GET("/index", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := engine.Group("/v1")
	{
		v1.GET("/", func(c *myweb.Context) {
			c.HTML(http.StatusOK, "<h1>hello myweb</h1>")
		})
		v1.GET("/hello", func(c *myweb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := engine.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *myweb.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *myweb.Context) {
			c.JSON(http.StatusOK, myweb.H{
				"username": c.PostForm("username"),
				"passowrd": c.PostForm("password"),
			})
		})
	}

	engine.Run(":9999")
}
