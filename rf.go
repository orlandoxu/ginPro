package ginPro

import "sync"

const (
	debugCode = iota
	releaseCode
	testCode
)

func New() *engine {
	return &engine{
		ctxPool: sync.Pool{
			New: func() interface{} {
				return &Context{}
			},
		},
	}
}
