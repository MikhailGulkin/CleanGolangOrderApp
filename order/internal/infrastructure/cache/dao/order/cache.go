package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/dao"
	base "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache/dao"
	"github.com/google/uuid"
)

type CacheDAOImpl struct {
	dao.OrderCacheDAO
	base.BaseRedisDAO
}

func (dao *CacheDAOImpl) GetOrder(orderID uuid.UUID) dto.Order {
	keys, _ := dao.Client.Keys(context.Background(), fmt.Sprintf("order:*:%s", orderID)).Result()
	for _, key := range keys {
		orderData, err := dao.Client.Get(context.Background(), key).Result()
		if err != nil {
			break
		}
		view := dto.Order{}

		err = json.Unmarshal([]byte(orderData), &view)
		if err != nil {
			break
		}
		return view //nolint
	}
	return dto.Order{}
}
func (dao *CacheDAOImpl) SaveOrder(order dto.Order) error {
	marshal, err := json.Marshal(order)
	if err != nil {
		return err
	}
	status := dao.Client.Set(
		context.Background(),
		fmt.Sprintf("order:%s:%s", order.Client.ClientID, order.OrderID),
		marshal,
		0,
	)
	return status.Err()
}
func (dao *CacheDAOImpl) DeleteOrder(orderID uuid.UUID) error {
	keys, _ := dao.Client.Keys(context.Background(), fmt.Sprintf("order:*:%s", orderID)).Result()
	for _, key := range keys {
		_, err := dao.Client.Del(context.Background(), key).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
