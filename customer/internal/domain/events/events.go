package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
)

const (
	CustomerCreated    = "CUSTOMER_CREATED"
	TransactionUpdated = "TRANSACTION_UPDATED"
	BalanceUpdated     = "BALANCE_UPDATED"
)

type CustomerCreatedEvent struct {
	FullName  vo.FullName        `json:"fullName"`
	AddressID uuid.UUID          `json:"addressID,omitempty"`
	Balance   vo.CustomerBalance `json:"balance"`
}

func NewCustomerCreatedEvent(
	aggregate common.Aggregate,
	fullName vo.FullName, addressID uuid.UUID,
	balance vo.CustomerBalance,
) (common.Event, error) {
	eventData := CustomerCreatedEvent{
		FullName:  fullName,
		AddressID: addressID,
		Balance:   balance,
	}
	baseEvent := common.NewBaseEvent(aggregate, CustomerCreated)
	if err := baseEvent.SetJsonData(&eventData); err != nil {
		return common.Event{}, err
	}
	return baseEvent, nil
}

type TransactionsUpdatedEvent struct {
	Transaction *entities.CustomerTransactions `json:"transaction"`
}

func NewTransactionsUpdatedEvent(
	aggregate common.Aggregate,
	transaction *entities.CustomerTransactions,
) (common.Event, error) {
	eventData := TransactionsUpdatedEvent{Transaction: transaction}
	baseEvent := common.NewBaseEvent(aggregate, TransactionUpdated)
	if err := baseEvent.SetJsonData(&eventData); err != nil {
		return common.Event{}, err
	}
	return baseEvent, nil
}

type BalanceUpdatedEvent struct {
	Balance vo.CustomerBalance `json:"balance"`
}

func NewBalanceUpdatedEvent(
	aggregate common.Aggregate,
	balance vo.CustomerBalance,
) (common.Event, error) {
	eventData := BalanceUpdatedEvent{Balance: balance}
	baseEvent := common.NewBaseEvent(aggregate, BalanceUpdated)
	if err := baseEvent.SetJsonData(&eventData); err != nil {
		return common.Event{}, err
	}
	return baseEvent, nil
}
