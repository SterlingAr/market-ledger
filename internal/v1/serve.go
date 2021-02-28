package v1

import (
	"context"
	v1 "github.com/SterlingAr/market-ledger/api/proto/v1"
	"github.com/SterlingAr/market-ledger/internal/pkg/database"
	"github.com/go-pg/pg/v10"
	"github.com/google/logger"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type MarketLedgerServer struct {
	v1.UnimplementedMarketLedgerServer
}

var db *pg.DB

func Serve(cmd *cobra.Command, args []string) {
	setup()

	purgeDb, _ := cmd.Flags().GetBool("purge-db")

	if purgeDb {
		cleanDB()
		err := seedDB()
		if err != nil {
			logger.Fatal(err)
		}
	}

	ctx := context.Background()

	router := runtime.NewServeMux()

	go grpcServer()

	err := v1.RegisterMarketLedgerHandlerServer(ctx, router, MarketLedgerServer{})

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	port := viper.GetString("http.port")
	address := ":" + port
	logger.Infof("Starting http server %v", address)

	logger.Error(http.ListenAndServe(address, router))
}

func grpcServer() {
	port := viper.GetString("grpc.port")
	address := ":" + port

	lis, err := net.Listen("tcp", address)

	if err != nil {
		logger.Error(err)
	}

	s := grpc.NewServer()

	logger.Infof("Starting gRPC server %v", address)

	if err := s.Serve(lis); err != nil {
		logger.Errorf("failed to serve: %v", err)
	}
}

func setup() {
	ctx := context.Background()
	db = database.Connection(ctx, database.ConnectionParams{
		User:     viper.GetString("database.postgres.db_user"),
		Password: viper.GetString("database.postgres.db_password"),
		Host:     viper.GetString("database.postgres.host"),
		Port:     viper.GetString("database.postgres.port"),
		Database: viper.GetString("database.postgres.db_name"),
		Schema:   viper.GetString("database.postgres.schema"),
	})
}
