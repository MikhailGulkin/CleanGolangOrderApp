package vo

import "time"

type OrderDeleted struct {
	Deleted  bool       `json:"deleted"`
	DeleteAt *time.Time `json:"delete_at"`
}

func (OrderDeleted) Create() OrderDeleted {
	return OrderDeleted{Deleted: false, DeleteAt: nil}
}
