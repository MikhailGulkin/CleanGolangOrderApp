package cache

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
)

type OrderCacheImpl struct {
	cache.OrderCache
	dao.OrderCacheDAO
}

func (o *OrderCacheImpl) OrderEvent(event interface{}) {
	switch e := event.(type) {
	case cache.OrderCreateSubscribe:
		order := o.OrderCacheDAO.GetOrder(e.OrderID)
		order.OrderStatus = e.OrderStatus
		order.OrderID = e.OrderID
		order.CreatedAt = e.CreatedAt
		order.Address = e.DeliveryAddress.FullAddress
		order.TotalPrice = e.TotalPrice
		order.SerialNumber = e.SerialNumber
		order.PaymentMethod = e.PaymentMethod
		order.Client = dto.Client{
			ClientID:   e.Client.ClientID,
			ClientName: e.Client.Username,
		}
		err := o.OrderCacheDAO.SaveOrder(order)
		if err != nil {
			return
		}
	case cache.OrderAddProductSubscribe:
		order := o.OrderCacheDAO.GetOrder(e.OrderID)
		order.OrderID = e.OrderID
		order.Products = append(order.Products, dto.Product{
			ProductID:   e.ProductID,
			Name:        e.ProductName,
			ActualPrice: e.ProductPrice,
		})
		order.Client = dto.Client{
			ClientID: e.ClientID,
		}
		err := o.OrderCacheDAO.SaveOrder(order)
		if err != nil {
			return
		}
	}
}
