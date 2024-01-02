package dispatcher

import(
	"fmt"
	"github.com/begopher/event"
)

func New(events ...int) event.Dispatcher {
	queues := make(map[int]map[string]event.Registration, len(events))
	for _, id := range events{
		queues[id] = make(map[string]event.Registration)
	}
	return dispacher{queues}
}

type dispacher struct {
	queues map[int]map[string]event.Registration
}

func (d dispacher) Send(event int) error {
	if queue, ok := d.queues[event]; ok {
		for _, registration := range queue {
			registration.Occur()
		}
		return nil
	}
	return fmt.Errorf("Event does not exist")
}

func (d dispacher) Bind(event int, reg event.Registration) error {
	if reg == nil {
		return fmt.Errorf("nil registration")
	}
	if queue, ok := d.queues[event]; ok {
		queue[reg.Name()] = reg
		return nil
	}
	return fmt.Errorf("Event does not exist")
}

func (d dispacher) Unbind(event int, reg event.Registration) error {
	if reg == nil {
		return fmt.Errorf("nil registration")
	}
	if queue, ok := d.queues[event]; ok {
		delete(queue, reg.Name())
		return nil
	}
	return fmt.Errorf("Event does not exist")

}
func (d dispacher) Publish(id int) bool {
	if _, ok := d.queues[id]; !ok {
		d.queues[id] = make(map[string]event.Registration)
	}
	return true
}
func (d dispacher) Unpublish(id int) bool {
	delete(d.queues, id)
	return true
}
