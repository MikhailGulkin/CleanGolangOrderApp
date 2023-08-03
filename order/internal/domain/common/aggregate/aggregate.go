package aggregate

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/events"

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
	length := len(r.GetEvents())
	cleared := make([]events.Event, length)
	copy(cleared, r.GetEvents())
	r.Events = r.Events[0:]
	return cleared
}
