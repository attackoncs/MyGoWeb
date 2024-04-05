package main

import (
	"fmt"
	"html/template"
	"myweb"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	engine := myweb.New()
	engine.Use(myweb.Logger())
	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./static")

	stu1 := &student{
		Name: "attackoncs", Age: 19,
	}
	stu2 := &student{
		Name: "MikeJordan", Age: 19,
	}
	engine.GET("/", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	engine.GET("/students", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", myweb.H{
			"title":  "myweb",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	engine.GET("/date", func(c *myweb.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", myweb.H{
			"title": "myweb",
			"now":   time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
		})
	})
	engine.Run(":9999")
}
