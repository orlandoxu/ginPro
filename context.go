package rf

import (
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
	//index    int8
	fullPath string

	// 洋葱
	Handers     HandlersChain
	HanderIndex int8
}

func (c *Context) Next() {
	//c.HanderIndex++
	//c.Handers[c.HanderIndex](c)

	c.HanderIndex++
	for c.HanderIndex < int8(len(c.Handers)) {
		c.Handers[c.HanderIndex](c)
		c.HanderIndex++
	}
}

func (c *Context) Abort() {
	c.HanderIndex = math.MaxInt8 / 2
}
