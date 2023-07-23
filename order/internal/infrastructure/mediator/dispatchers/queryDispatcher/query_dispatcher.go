package querydispatcher

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/query"
	"reflect"
)

type QueryDispatcher interface {
	Query(query interface{}) (interface{}, error)
	RegisterQueryHandler(query interface{}, handler QueryHandler)
}

type QueryDispatcherImpl struct {
	QueryHandlers map[reflect.Type]QueryHandler
}
type QueryHandler interface {
	Handle(query interface{}) (interface{}, error)
}

func (c *QueryDispatcherImpl) Query(query interface{}) (interface{}, error) {
	var handler QueryHandler
	c.GetQueryHandler(query, &handler)
	return handler.Handle(query)
}
func (c *QueryDispatcherImpl) GetQueryHandler(query interface{}, receiver interface{}) {
	reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(c.QueryHandlers[reflect.TypeOf(query)]))
}

type CreateProductQueryHandler struct {
	query.GetAllProducts
}

func (c *QueryDispatcherImpl) RegisterQueryHandler(query interface{}, handler QueryHandler) {
	c.QueryHandlers[reflect.ValueOf(query).Type()] = handler
}
func (c *CreateProductQueryHandler) Query(q interface{}) (interface{}, error) {
	return c.GetAllProducts.Get(q.(query.GetAllProductsQuery))
}
