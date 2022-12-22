package ginp

import (
	"encoding/json"
	"math"
	"net/http"
)

// 上下文
type Context struct {
	Method   string
	Path     string
	RawQuery string
	Host     string
	Request  *http.Request
	Writer   *http.ResponseWriter
	//index    int8
	fullPath string

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
