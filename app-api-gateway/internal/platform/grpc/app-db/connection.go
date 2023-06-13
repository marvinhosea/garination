package app_db

import (
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDatabaseConn(db config.AppDb) (proto.DatabaseServiceClient, error) {
	conn, err := grpc.Dial(db.Host+":"+db.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewDatabaseServiceClient(conn), nil
}
