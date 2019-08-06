package mx

type HybridSequential struct {

}

func (hs *HybridSequential) Add(b... HybridBlock) error {
	return nil
}

func (hs *HybridSequential) Hybridize() error {
	return nil
}

func (hs *HybridSequential) Forward(F *Function, x Symbol) error {
	return nil
}


