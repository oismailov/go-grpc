package main

import (
	"log"

	"github.com/abhiyerra/landingcrew-cli/cmd"
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
