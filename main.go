package main

import (
	"fmt"
	"myweb"
	"net/http"
)

func main() {
	engine := myweb.New()
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	})
	engine.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q]=%q\n", k, v)
		}
	})
	engine.Run(":9999")
}
