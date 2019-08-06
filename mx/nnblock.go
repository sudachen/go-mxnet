package mx

type Function struct {}
type Symbol struct {}
type Initializer interface{}

type InitParams struct {
	Init Initializer
	Ctx  Context
	Verbose bool
	Force bool
}

type HybridBlock interface {
	Hybridize() error
	Forward(F *Function, x Symbol) error
	Initialize(params InitParams)
}


