package event

type Dispatcher interface {
	Send(int) error
	Bind(int, Registration) error
	Unbind(int, Registration) error
}

