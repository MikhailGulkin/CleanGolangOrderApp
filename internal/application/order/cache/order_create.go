package cache

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"time"
)

type OrderCacheImpl struct {
	cache.OrderCache
}

func (o *OrderCacheImpl) OrderCreate(event interface{}) {
	switch event.(type) {
	case cache.OrderCreateSubscribe:
		time.Sleep(time.Second * 2)
		fmt.Println("OrderCreateSubscribe", event)
	case cache.OrderAddProductSubscribe:
		time.Sleep(time.Second * 2)
		fmt.Println("OrderAddProductSubscribe", event)
	}
}
