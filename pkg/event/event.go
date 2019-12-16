package event

type Interface interface {
	Handle() error
}
