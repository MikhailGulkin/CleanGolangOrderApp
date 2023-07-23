package order

import (
	"context"
	"encoding/json"
	"fmt"
	filter "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/reader"
	r "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache/reader"
	"github.com/google/uuid"
	"sort"
)

type CacheReaderImpl struct {
	reader.OrderCacheReader
	r.BaseRedisReader
}

func (reader *CacheReaderImpl) GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error) {
	keys, err := reader.Client.Keys(context.Background(), fmt.Sprintf("order:*:*")).Result()
	if err != nil {
		return []dto.Order{}, err
	}
	l := uint(len(keys))
	if l == 0 {
		return []dto.Order{}, nil
	}
	if filters.Offset >= l {
		filters.Offset = l - 1
	}
	end := filters.Offset + filters.Limit
	if end > l {
		end = l
	}
	keys = keys[filters.Offset:end]
	orders := make([]dto.Order, len(keys))
	for index, key := range keys {
		orderData, err := reader.Client.Get(context.Background(), key).Result()
		if err != nil {
			continue
		}
		order := dto.Order{}

		err = json.Unmarshal([]byte(orderData), &order)
		if err != nil {
			continue
		}
		orders[index] = order
	}
	sort.SliceStable(orders, func(i, j int) bool {
		if filters.Order == filter.ASC {
			return orders[i].TotalPrice < orders[j].TotalPrice
		}
		return orders[i].TotalPrice > orders[j].TotalPrice
	})
	return orders, nil
}
func (reader *CacheReaderImpl) GetAllOrdersByUserID(userID uuid.UUID, filters filters.GetAllOrdersByUserIDFilters) ([]dto.Order, error) {
	keys, err := reader.Client.Keys(context.Background(), fmt.Sprintf("order:%s:*", userID)).Result()
	if err != nil {
		return []dto.Order{}, err
	}
	l := uint(len(keys))
	if l == 0 {
		return []dto.Order{}, nil
	}
	if filters.Offset >= l {
		filters.Offset = l - 1
	}
	end := filters.Offset + filters.Limit
	if end > l {
		end = l
	}
	keys = keys[filters.Offset:end]
	orders := make([]dto.Order, len(keys))
	for index, key := range keys {
		orderData, err := reader.Client.Get(context.Background(), key).Result()
		if err != nil {
			continue
		}
		order := dto.Order{}

		err = json.Unmarshal([]byte(orderData), &order)
		if err != nil {
			continue
		}
		orders[index] = order
	}
	sort.SliceStable(orders, func(i, j int) bool {
		if filters.Order == filter.ASC {
			return orders[i].TotalPrice < orders[j].TotalPrice
		}
		return orders[i].TotalPrice > orders[j].TotalPrice
	})
	return orders, nil
}

func (reader *CacheReaderImpl) GetOrderByID(orderID uuid.UUID) (dto.Order, error) {
	keys, err := reader.Client.Keys(context.Background(), fmt.Sprintf("order:*:%s", orderID)).Result()
	if err != nil {
		return dto.Order{}, err
	}
	for _, key := range keys {
		orderData, err := reader.Client.Get(context.Background(), key).Result()
		if err != nil {
			return dto.Order{}, err
		}
		view := dto.Order{}

		err = json.Unmarshal([]byte(orderData), &view)
		if err != nil {
			return dto.Order{}, err
		}
		return view, nil //nolint
	}
	customErr := exceptions.OrderIDNotExist{}.Exception(orderID.String())
	return dto.Order{}, &customErr
}
