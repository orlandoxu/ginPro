package ginp

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
)

// 上下文
type Context struct {
	Method string
	Path   string
	Url    *url.URL
	//RawQuery string
	Host    string
	Request *http.Request
	Writer  *http.ResponseWriter
	//index    int8
	fullPath   string
	Body       []byte
	contextMap map[string]interface{} // 可以记录上下文

	// 洋葱
	Handers     HandlersChain
	HanderIndex int8
}

func (c *Context) Next() {
	c.HanderIndex++
	for c.HanderIndex < int8(len(c.Handers)) {
		c.Handers[c.HanderIndex](c)
		c.HanderIndex++
	}
}

func (c *Context) Abort() {
	c.HanderIndex = math.MaxInt8 / 2
}

func (c *Context) Json(ret int, params ...interface{}) {
	h := make(map[string]interface{})
	h["ret"] = ret

	if len(params) == 1 {
		h["data"] = params[0]
	} else {
		h["msg"] = params[0]
		h["data"] = params[1]
	}

	str, err := json.Marshal(h)
	if err != nil {
		// 如何处理呢？
	}

	(*c.Writer).Write(str)
}

func (c *Context) SendAbort(ret int, params ...interface{}) {
	c.Json(ret, params...)
	c.Abort()
}

// 兼容以前的代码
func (c *Context) Send(ret int, params ...interface{}) {
	c.Json(ret, params...)
}

func (c *Context) BindBodyData(t any) error {
	if c.Body != nil {
		return json.Unmarshal(c.Body, t)
	}

	return nil
}

func (c *Context) Set(key string, value interface{}) {
	if c.contextMap == nil {
		c.contextMap = make(map[string]interface{})
	}

	c.contextMap[key] = value
}

func (c *Context) Get(key string) (interface{}, bool) {
	r, isOk := c.contextMap[key]

	return r, isOk
}

func (c *Context) GetString(key string) string {
	str, _ := c.Get(key)

	return str.(string)
}

func (c *Context) Query(key string) string {
	values := c.Url.Query()

	k, isOk := values[key]

	if !isOk || len(k) == 0 {
		return ""
	}

	return k[0]
}
