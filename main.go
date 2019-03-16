package main

import (
	"log"

	"github.com/oismailov/go-grpc/cmd"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cmd.Execute(conn)
}
