package db

import (
	"context"
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/jackc/pgx/v4"
)

type EventStore struct {
	Conn Connection
}

func NewEventStore(conn Connection) persistence.EventStore {
	return &EventStore{
		Conn: conn,
	}
}

func (es *EventStore) Create(customer common.Aggregate, tx interface{}) error {
	conn := tx.(pgx.Tx)
	query := `INSERT INTO public.entities (entity_id, entity_type, entity_version) 
				VALUES ($1, $2, $3);`
	entity := ConvertAggregateToEntity(customer)
	_, err := conn.Exec(
		context.Background(),
		query, entity.EntityID, entity.EventType, entity.EntityVersion,
	)
	if err != nil {
		return err
	}
	query = `INSERT INTO public.events (event_id, entity_id, event_type, entity_type, event_data, created_at)
				VALUES ($1, $2, $3, $4, $5, $6);`
	for _, e := range customer.GetUncommittedEvents() {
		event := ConvertDomainEventToEventModel(e)
		_, err = conn.Exec(
			context.Background(),
			query, event.EventID, event.EntityID, event.EventType, event.EntityType, event.EventData, event.CreatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
func (es *EventStore) Exists(id string) error {
	query := "SELECT entity_id FROM public.entities WHERE entity_id = $1"
	row := es.Conn.QueryRow(context.Background(), query, id)
	var customerID string
	err := row.Scan(&customerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if customerID == id {
		return application.ErrCustomerAlreadyExists
	}
	return nil
}
func (es *EventStore) Update(_ common.Aggregate, _ interface{}) error {
	return nil
}
func (es *EventStore) Find(_ string) (common.Aggregate, error) {
	return nil, nil
}
