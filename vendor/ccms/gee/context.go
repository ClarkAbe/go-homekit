package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
	// engine pointer
	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:   req.URL.Path,
		Method: req.Method,
		Req:    req,
		Writer: w,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) AnyForm(key string) string {
	c.Req.ParseForm()
	return strings.Join(c.Req.Form[key], "-")
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) QueInt(key string) int {
	k, _ := strconv.Atoi(c.Req.URL.Query().Get(key))
	return k
}

func (c *Context) QueInt64(key string) int64 {
	k, _ := strconv.ParseInt(c.Req.URL.Query().Get(key), 10, 64)
	return k
}

func (c *Context) Cookie(key string) string {
	if v,e := c.Req.Cookie(key);e == nil {
		return v.Value
	}
	return ""
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML template render
// refer https://golang.org/pkg/html/template/
func (c *Context) Template(code int, name string, data interface{}) {
	//c.Writer.Header().Set("Content-Type", "text/html")
	c.Status(code)
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context) HTML(code int, name string) {
	//fmt.Println(name)
	//c.Writer.Header().Set("Content-Type", "text/html")
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	if _, ok := c.engine.htmlFiles[name]; ok == true {
		c.Status(code)
		c.Writer.Write(c.engine.htmlFiles[name])
	}else{
		c.Fail(500, "HTML Is Not Found!")
	}
}