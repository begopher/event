package event

type Dispatcher interface {
	Send(int) error
	Publish(int) bool
	Unpublish(int) bool
	Bind(int, Registration) error
	Unbind(int, Registration) error
}

