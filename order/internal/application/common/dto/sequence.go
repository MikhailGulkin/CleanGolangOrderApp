package dto

import f "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/filters"

type BaseSequence struct {
	Count  uint        `json:"count"`
	Limit  uint        `json:"limit,omitempty"`
	Offset uint        `json:"offset,omitempty"`
	Order  f.BaseOrder `json:"order,omitempty"`
}
