package ledger

import (
	"context"
	"github.com/SterlingAr/market-ledger/internal/pkg/database"
	"github.com/google/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	db = database.Connection(ctx, database.ConnectionParams{
		User:     "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     "5432",
		Database: "market",
		Schema:   "ledger",
	})
	err := database.CreateSchema(db, "ledger")
	if err != nil {
		logger.Fatal(err)
	}
	cleanDB()
	os.Exit(m.Run())
}
