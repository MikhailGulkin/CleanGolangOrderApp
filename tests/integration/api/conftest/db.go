package conftest

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
	"testing"
	"time"
)

func CleanTables(conn *gorm.DB) {
	var tables []string
	if err := conn.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	tx := conn.Begin()
	for _, table := range tables {
		tx.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", table))
	}
	tx.Commit()
}

type TestDatabase struct {
	instance testcontainers.Container
}

func NewTestDatabase(t *testing.T) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432"},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "OrderAppTest",
		},
		WaitingFor: wait.ForListeningPort("5432"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	_, _, _ = postgres.Exec(context.Background(), []string{"CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"})
	return &TestDatabase{
		instance: postgres,
	}
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "5432")
	require.NoError(t, err)
	return p.Int()
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://postgres:1234@localhost:%d/OrderAppTest", db.Port(t))
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}
