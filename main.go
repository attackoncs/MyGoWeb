package main

import (
	"log"
	"myweb"
	"net/http"
	"time"
)

func main() {
	engine := myweb.New()
	engine.Use(myweb.Logger())
	engine.GET("/", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Myweb</h1>")
	})
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
	v2.Use(onlyForV2())
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

func onlyForV2() myweb.HandlerFunc {
	return func(c *myweb.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
