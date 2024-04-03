package myweb

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试注册路由
func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

// 测试解析路由
func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name", "*"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

// 测试解析路由得到的路由参数
func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name shoud be equal to 'geektutu'")
	}

	fmt.Printf("matched path:%s,params['name']:%s\n", n.pattern, ps["name"])
}
