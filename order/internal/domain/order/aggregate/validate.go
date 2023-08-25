package aggregate

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"

func (o *Order) PreprocessOrder() error {
	if err := o.checkNotDeleted(); err != nil {
		return err
	}
	if err := o.checkNotClosed(); err != nil {
		return err
	}
	return nil
}
func (o *Order) checkNotDeleted() error {
	if o.Deleted {
		exception := exceptions.OrderIsDeleted{}.Exception(o.OrderID.ToString())
		return &exception
	}
	return nil
}

func (o *Order) checkNotClosed() error {
	if o.Closed {
		exception := exceptions.OrderIsClosed{}.Exception(o.OrderID.ToString())
		return &exception
	}
	return nil
}
