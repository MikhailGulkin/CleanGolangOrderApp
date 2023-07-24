package vo

import "time"

type OrderInfo struct {
	Date         time.Time
	SerialNumber int
	Closed       bool
}

func (OrderInfo) Create(serialNumber int) OrderInfo {
	return OrderInfo{SerialNumber: serialNumber, Date: time.Now(), Closed: false}
}
