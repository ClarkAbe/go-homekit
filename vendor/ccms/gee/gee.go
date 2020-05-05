package gee

import (
	"io/ioutil"
	"html/template"
	"net/http"
	"path"
	"time"
	"errors"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc // support middleware
		parent      *RouterGroup  // support nesting
		engine      *Engine       // all groups share a Engine instance
	}

	Engine struct {
		*RouterGroup
		http          *http.Server
		http_start    bool
		router        *router
		groups        []*RouterGroup     // store all groups
		htmlTemplates *template.Template // for html render
		htmlFiles     map[string][]byte // for html render
		funcMap       template.FuncMap   // for html render
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func NewUse(log bool, rec bool) *Engine {
	engine := New()
	if log {
		engine.Use(Logger())
	}else if rec {
		engine.Use(Recovery())
	}else if rec && log {
		engine.Use(Logger(), Recovery())
	}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp

	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//add Get and Post request 
func (group *RouterGroup) RESTFUL(pattern string, handler HandlerFunc) {
	s := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
	for i := 0; i < len(s); i++ {
		group.addRoute(s[i], pattern, handler)
	}
}

//Custom request
func (group *RouterGroup) Custom(method, pattern string, handler HandlerFunc) {
	group.addRoute(strings.ToUpper(method), pattern, handler)
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// create static handler
func (group *RouterGroup) createGzipStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := GzipHandler(http.StripPrefix(absolutePath, http.FileServer(fs)))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}


// serve static files
func (group *RouterGroup) GzipStatic(relativePath string, root string) {
	handler := group.createGzipStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

// for custom render function
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadTemplateGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlFiles = make(map[string][]byte)
	rd, _ := ioutil.ReadDir(pattern)
	for _, fi := range rd {
		if str, err := ioutil.ReadFile(pattern + "/" + fi.Name()); err == nil {
			engine.htmlFiles[fi.Name()] = str
		}
	}
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) RunTls(addr string, pem string, key string) (err error) {
	return http.ListenAndServeTLS(addr, pem, key, engine)
}

func (engine *Engine) RunSrv(addr string) (error) {
	if engine.http_start == true {
		return errors.New("http server is start!")
	}
	engine.http = &http.Server{Addr: addr, Handler: engine}
	var err error
	go func(){
		err = engine.http.ListenAndServe()
	}()
	time.Sleep(1 * time.Second) //等待1秒钟看看Goroutine有没有报错!
	if err == nil {
		engine.http_start = true 
	}
	return err
}

func (engine *Engine) ShutdownSrv() (error){
	if engine.http_start == false {
		return errors.New("http server not start!")
	}
	if err := engine.http.Shutdown(nil); err != nil {
		return err
	}
	engine.http_start = false
	return nil
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}
