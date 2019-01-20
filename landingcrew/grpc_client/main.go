package grpc_client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:6000"
)

func GetConnectionSetting() (*grpc.ClientConn, context.Context, context.CancelFunc) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn, ctx, cancel
}
