package rf

const (
	debugCode = iota
	releaseCode
	testCode
)

func New() *engine {
	engine := &engine{}

	return engine
}
