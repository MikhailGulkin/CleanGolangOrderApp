package commands

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/stretchr/testify/mock"
)

type MockEventStore struct {
	mock.Mock
}

func (t *MockEventStore) Create(aggregate common.Aggregate, tx interface{}) error {
	args := t.Called(aggregate, tx)
	return args.Error(0)
}
func (t *MockEventStore) Update(aggregate common.Aggregate, tx interface{}) error {
	args := t.Called(aggregate, tx)
	return args.Error(0)
}
func (t *MockEventStore) Load(aggregate common.Aggregate) error {
	args := t.Called(aggregate)
	return args.Error(0)
}

func (t *MockEventStore) Exists(id string) error {
	args := t.Called(id)
	return args.Error(0)
}

type TestOutbox struct {
}

func (t *TestOutbox) AddEvents(_ []common.Event, _ interface{}) error {
	return nil
}

type TestUoW struct {
}

func (t *TestUoW) Begin() (interface{}, error) {
	return nil, nil
}
func (t *TestUoW) Commit() error {
	return nil
}
func (t *TestUoW) Rollback() error {
	return nil
}

type TestUoWManager struct {
}

func (t *TestUoWManager) GetUoW() persistence.UoW {
	return &TestUoW{}
}
