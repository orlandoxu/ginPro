package rf

import "regexp"

// -- 路由模块
type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

var middlewareChain []HandlerFunc
var router map[string]HandlersChain // TODO：没做锁，需要考虑是否需要加锁

func init() {
	router = make(map[string]HandlersChain)
	middlewareChain = []HandlerFunc{}
}

func PushToChain(url string, handlerFunc ...HandlerFunc) {

}

// TODO: 这儿可以提供一种，根据正则表达式来匹配router的逻辑
func PushMiddlewareToChain(reg string, handlerFunc ...HandlerFunc) {
	middlewareChain = append(middlewareChain, handlerFunc...)

	for k, v := range router {
		isOk, _ := regexp.MatchString(reg, k)
		if reg == "" || isOk {
			router[k] = append(v, handlerFunc...)
		}
	}
}

func GetHandlersChain(url string) HandlersChain {
	r, has := router[url]
	if !has {
		return nil
	}

	return r
}
