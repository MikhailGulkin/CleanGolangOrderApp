package db

import (
	"context"
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/exceptions"
	"github.com/jackc/pgx/v4"
)

type EventStore struct {
	Conn Connection
}

func NewEventStore(conn Connection) aggregate.EventStore {
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
		return exceptions.ErrCustomerAlreadyExists
	}
	return nil
}
func (es *EventStore) Update(aggregate common.Aggregate, tx interface{}) error {
	conn := tx.(pgx.Tx)
	query := `UPDATE public.entities SET entity_version = $1 WHERE entity_id = $2;`
	entity := ConvertAggregateToEntity(aggregate)
	_, err := conn.Exec(
		context.Background(),
		query,
		entity.EntityVersion, entity.EntityID,
	)
	if err != nil {
		return err
	}
	query = `INSERT INTO public.events (event_id, entity_id, event_type, entity_type, event_data, created_at)
				VALUES ($1, $2, $3, $4, $5, $6);`
	for _, e := range aggregate.GetUncommittedEvents() {
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
func (es *EventStore) Load(aggregate common.Aggregate) error {
	query := `SELECT event_id, entity_id, event_type, entity_type, event_data, created_at
				FROM public.events WHERE entity_id = $1;`
	rows, err := es.Conn.Query(context.Background(), query, aggregate.GetID())
	if err != nil {
		return err
	}
	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.EventID, &event.EntityID, &event.EventType, &event.EntityType, &event.EventData, &event.CreatedAt)
		if err != nil {
			return err
		}
		events = append(events, event)
	}
	err = aggregate.Load(ConvertEventsToDomainEvents(events))
	if err != nil {
		return err
	}
	return nil
}

type Outbox struct {
	Conn Connection
}

func NewOutbox(conn Connection) persistence.Outbox {
	return &Outbox{
		Conn: conn,
	}
}

func (o *Outbox) AddEvents(events []common.Event, tx interface{}) error {
	conn := tx.(pgx.Tx)
	query := `INSERT INTO public.outbox (exchange, route, payload, aggregate_id)
				VALUES ($1, $2, $3, $4);`
	for _, e := range events {
		event := ConvertDomainEventToOutboxMessage(e)
		_, err := conn.Exec(
			context.Background(),
			query, event.Exchange, event.Route, event.Payload, event.AggregateID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
