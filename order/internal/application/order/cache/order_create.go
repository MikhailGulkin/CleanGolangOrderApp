package cache

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
)

type OrderCacheImpl struct {
	cache.OrderCache
	dao.OrderCacheDAO
	logger.Logger
}

func (o *OrderCacheImpl) OrderCreateEvent(event cache.OrderCreateSubscribe) {
	order := o.OrderCacheDAO.GetOrder(event.OrderID)
	order.OrderStatus = event.OrderStatus
	order.OrderID = event.OrderID
	order.CreatedAt = event.CreatedAt
	order.Address = event.DeliveryAddress.FullAddress
	order.TotalPrice = event.TotalPrice
	order.SerialNumber = event.SerialNumber
	order.PaymentMethod = event.PaymentMethod
	order.Client = dto.Client{
		ClientID:   event.Client.ClientID,
		ClientName: event.Client.Username,
	}
	err := o.OrderCacheDAO.SaveOrder(order)
	if err != nil {
		return
	}
}
func (o *OrderCacheImpl) OrderAddProductEvent(event cache.OrderAddProductSubscribe) {
	order := o.OrderCacheDAO.GetOrder(event.OrderID)
	order.OrderID = event.OrderID
	order.Products = append(order.Products, dto.Product{
		ProductID:   event.ProductID,
		Name:        event.ProductName,
		ActualPrice: event.ProductPrice,
	})
	order.Client = dto.Client{
		ClientID: event.ClientID,
	}
	err := o.OrderCacheDAO.SaveOrder(order)
	if err != nil {
		return
	}
}
func (o *OrderCacheImpl) OrderDeleteEvent(event cache.OrderDeleteEvent) {
	err := o.OrderCacheDAO.DeleteOrder(event.OrderID)
	if err != nil {
		o.Logger.Info(fmt.Sprintf("Canno't delete order from cache id %s", event.OrderID))
	}
}
