package common

type AggregateType string

const (
	aggregateStartVersion                = -1
	aggregateAppliedEventsInitialCap     = 10
	aggregateUncommittedEventsInitialCap = 10
)

type When interface {
	When(event Event) error
}

// Apply process Aggregate Event
type Apply interface {
	Apply(event Event) error
}

// Load create Aggregate state from Event's.
type Load interface {
	Load(events []Event) error
}

type when func(event Event) error
type Aggregate interface {
	When
	AggregateRoot
}
type AggregateRoot interface {
	GetUncommittedEvents() []Event
	GetID() string
	SetID(id string) *AggregateBase
	GetVersion() int64
	SetType(aggregateType AggregateType)
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	RaiseEvent(event Event) error
	String() string
	Load
	Apply
}

type AggregateBase struct {
	ID                string
	Version           int64
	Type              AggregateType
	AppliedEvents     []Event
	UncommittedEvents []Event
	when              when
}

// NewAggregateBase create new AggregateBase
// main purpose of this function is to set when function
func NewAggregateBase(when when) *AggregateBase {
	if when == nil {
		return nil
	}

	return &AggregateBase{
		Version:           aggregateStartVersion,
		AppliedEvents:     make([]Event, 0, aggregateAppliedEventsInitialCap),
		UncommittedEvents: make([]Event, 0, aggregateUncommittedEventsInitialCap),
		when:              when,
	}
}
