package rf

type engine struct{}

func (e *engine) Post(url string, handlerFunc ...HandlerFunc) {

}

func (e *engine) Use(handlerFunc ...HandlerFunc) {
	PushMiddlewareToChain("", handlerFunc...)
}

func (e *engine) UseByRegex(reg string, handlerFunc ...HandlerFunc) {
	PushMiddlewareToChain(reg, handlerFunc...)
}
