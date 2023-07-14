package dto

import f "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"

type BaseSequence struct {
	Count  uint        `json:"count"`
	Limit  uint        `json:"limit,omitempty"`
	Offset uint        `json:"offset,omitempty"`
	Order  f.BaseOrder `json:"order,omitempty"`
}
