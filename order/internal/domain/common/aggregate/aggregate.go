package aggregate

import "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/events"

type AggregateRoot struct {
	Events []events.Event
}

func (r *AggregateRoot) RecordEvent(event events.Event) {
	r.Events = append(r.Events, event)
}
func (r *AggregateRoot) GetEvents() []events.Event {
	return r.Events
}
func (r *AggregateRoot) PullEvents() []events.Event {
	length := len(r.Events)
	cleared := make([]events.Event, length)
	copy(cleared, r.Events)
	r.Events = r.Events[0:]
	return cleared
}
