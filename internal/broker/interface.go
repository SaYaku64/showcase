package broker

type IBroker interface {
	Run()
	Send(address string, data any)
}
