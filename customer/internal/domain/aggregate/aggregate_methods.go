package aggregate

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

func (a *CustomerAggregate) When(evt common.Event) error {

	switch evt.GetEventType() {
	case v1.OrderCreated:
		return a.onOrderCreated(evt)
	default:
		return common.ErrInvalidEventType
	}
}
func (a *CustomerAggregate) onOrderCreated(evt common.Event) error {
	var eventData v1.CustomerCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return fmt.Errorf("get Json Data, err: %w", err)
	}
}
