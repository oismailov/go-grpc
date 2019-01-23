package main

import (
	"github.com/abhiyerra/landingcrew-cli/landingcrew/cmd"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:6000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cmd.Execute(conn)
}
