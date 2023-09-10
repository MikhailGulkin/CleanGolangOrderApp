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
	AvatarUriCreated   = "AVATAR_URI_CREATED"
)

type CustomerCreatedEvent struct {
	ID        string      `json:"id,omitempty"`
	FullName  vo.FullName `json:"fullName"`
	AddressID uuid.UUID   `json:"addressID,omitempty"`
	Balance   vo.Balance  `json:"balance"`
	Email     vo.Email    `json:"email"`
}

type AvatarUriCreatedEvent struct {
	Uri string `json:"uri, omitempty"`
}

// NewCustomerCreatedEvent creates new CustomerCreatedEvent
func NewCustomerCreatedEvent(
	aggregate common.Aggregate,
	fullName vo.FullName,
	addressID uuid.UUID,
	email vo.Email,
	balance vo.Balance,
) (common.Event, error) {
	eventData := CustomerCreatedEvent{
		ID:        aggregate.GetID(),
		Email:     email,
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

// NewAvatarUriCreatedEvent creates new AvatarUriCreatedEvent
func NewAvatarUriCreatedEvent(aggregate common.Aggregate) (common.Event, error) {
	eventData := AvatarUriCreatedEvent{
		Uri: uuid.New().String(),
	}
	baseEvent := common.NewBaseEvent(aggregate, AvatarUriCreated)
	if err := baseEvent.SetJsonData(&eventData); err != nil {
		return common.Event{}, err
	}
	return baseEvent, nil
}

type TransactionsUpdatedEvent struct {
	Transaction *entities.CustomerTransactions `json:"transaction"`
}

// NewTransactionsUpdatedEvent creates new TransactionsUpdatedEvent
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
	Balance vo.Balance `json:"balance"`
}

// NewBalanceUpdatedEvent creates new BalanceUpdatedEvent
func NewBalanceUpdatedEvent(
	aggregate common.Aggregate,
	balance vo.Balance,
) (common.Event, error) {
	eventData := BalanceUpdatedEvent{Balance: balance}
	baseEvent := common.NewBaseEvent(aggregate, BalanceUpdated)
	if err := baseEvent.SetJsonData(&eventData); err != nil {
		return common.Event{}, err
	}
	return baseEvent, nil
}
