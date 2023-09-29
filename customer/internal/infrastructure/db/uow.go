package db

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/persistence"
	"github.com/jackc/pgx/v4"
)

type UoW struct {
	Conn Connection
	Tx   pgx.Tx
}

// Commit commits the transaction.
func (u *UoW) Commit() error {
	return u.Tx.Commit(context.Background())
}

// Rollback aborts the transaction.
func (u *UoW) Rollback() error {
	err := u.Tx.Rollback(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Begin starts a transaction.
func (u *UoW) Begin() (interface{}, error) {
	tx, err := u.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	u.Tx = tx
	return u.Tx, nil
}

type UoWManager struct {
	Conn Connection
}

// GetUoW returns a new UoW.
func (u *UoWManager) GetUoW() persistence.UoW {
	return &UoW{
		Conn: u.Conn,
	}
}

func NewUoWManager(conn Connection) *UoWManager {
	return &UoWManager{
		Conn: conn,
	}
}
