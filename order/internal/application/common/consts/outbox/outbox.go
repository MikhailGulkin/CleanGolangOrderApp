package outbox

type EventStatus uint

const (
	Undefined EventStatus = iota
	Awaiting              = 1
	Processed             = 2
	Sagas                 = 3
	Rejected              = 4
)
