package main

import (
	"myweb"
	"net/http"
)

func main() {
	engine := myweb.New()
	engine.GET("/", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "<h1>hello myweb</h1>")
	})
	engine.GET("/hello", func(c *myweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	engine.GET("/hello/:name", func(c *myweb.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	})
	engine.GET("/assets/*filepath", func(c *myweb.Context) {
		c.JSON(http.StatusOK, myweb.H{"filepath": c.Param("filepath")})
	})
	engine.POST("/login", func(c *myweb.Context) {
		c.JSON(http.StatusOK, myweb.H{
			"username": c.PostForm("username"),
			"passowrd": c.PostForm("password"),
		})
	})
	engine.Run(":9999")
}
