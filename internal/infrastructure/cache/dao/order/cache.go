package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	base "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/dao"
	"github.com/google/uuid"
)

type CacheDAOImpl struct {
	dao.OrderCacheDAO
	base.BaseRedisDAO
}

func (dao *CacheDAOImpl) GetOrder(orderID uuid.UUID) cache.OrderEvent {
	keys, _ := dao.Client.Keys(context.Background(), fmt.Sprintf("order:*:%s", orderID)).Result()
	for _, key := range keys {
		orderData, err := dao.Client.Get(context.Background(), key).Result()
		if err != nil {
			break
		}
		event := cache.OrderEvent{}

		err = json.Unmarshal([]byte(orderData), &event)
		if err != nil {
			break
		}
		return event //nolint
	}
	return cache.OrderEvent{}
}
func (dao *CacheDAOImpl) SaveOrder(event cache.OrderEvent) error {
	marshal, err := json.Marshal(event)
	if err != nil {
		return err
	}
	status := dao.Client.Set(
		context.Background(),
		fmt.Sprintf("order:%s:%s", event.Client.ClientID, event.OrderID),
		marshal,
		0,
	)
	return status.Err()
}
