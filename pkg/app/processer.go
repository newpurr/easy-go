package app

type Processer interface {
	MustStart()
	Stop()
}
