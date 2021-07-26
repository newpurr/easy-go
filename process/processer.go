package process

type Processer interface {
	MustStart()
	Stop()
}
