package rf

import (
	"log"
	"net/http"
	"sync"
)

type engine struct {
	ctxPool sync.Pool
}

func (e *engine) Post(url string, handlerFunc ...HandlerFunc) {
	PushToChain(url, handlerFunc...)
}

func (e *engine) Use(handlerFunc ...HandlerFunc) {
	PushMiddlewareToChain("", handlerFunc...)
}

func (e *engine) UseByRegex(reg string, handlerFunc ...HandlerFunc) {
	PushMiddlewareToChain(reg, handlerFunc...)
}

// ServeHTTP conforms to the http.Handler interface.
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(w)
	log.Println(req)

	c := e.ctxPool.Get().(*Context)
	c.Path = req.URL.Path
	c.RawQuery = req.URL.RawQuery
	c.Host = req.Host
	//c.writermem.reset(w)
	//c.Request = req
	//c.reset()
	//
	e.handleHTTPRequest(c)
	//
	e.ctxPool.Put(c)
}

// 这儿是处理的业务逻辑
func (e *engine) handleHTTPRequest(c *Context) {

	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false
	_ = httpMethod
	_ = rPath
	_ = unescape

	// Find root of the tree for the given HTTP method
}

func (e *engine) Run(host string) (err error) {
	return http.ListenAndServe(host, e)
}
