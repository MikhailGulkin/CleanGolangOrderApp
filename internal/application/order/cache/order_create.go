package cache

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
)

type OrderCacheImpl struct {
	cache.OrderCache
}

func (o *OrderCacheImpl) OrderCreate(event interface{}) {
	switch event.(type) {
	case cache.OrderCreateSubscribe:
		fmt.Println("1234")
	}
}
