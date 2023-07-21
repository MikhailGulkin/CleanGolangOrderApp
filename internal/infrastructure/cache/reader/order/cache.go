package order

import (
	"context"
	"encoding/json"
	"fmt"
	filter "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	r "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/reader"
	"github.com/google/uuid"
	"sort"
)

type CacheReaderImpl struct {
	reader.OrderCacheReader
	r.BaseRedisReader
}

func (reader *CacheReaderImpl) GetAllOrdersByUserID(userID uuid.UUID, filters filters.GetAllOrdersByUserIDFilters) ([]dto.Order, error) {
	keys, err := reader.Client.Keys(context.Background(), fmt.Sprintf("order:%s:*", userID)).Result()
	if err != nil {
		return []dto.Order{}, err
	}
	l := uint(len(keys))
	if l == 0 {
		return []dto.Order{}, err
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