package lib

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the email/apikey
type Authentication struct {
	Email  string
	APIKey string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"email":  a.Email,
		"apikey": a.APIKey,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

// Connect establishes a grpc authenticated TLS connection to the go-grpc API server.
func Connect(target string) (*grpc.ClientConn, error) {
	// Setup the login/pass
	auth := Authentication{
		Email:  os.Getenv("go-grpc_EMAIL"),
		APIKey: os.Getenv("go-grpc_API_KEY"),
	}

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("../../cert/server-cert.pem", "go-grpc")
	if err != nil {
		return nil, fmt.Errorf("could not load tls cert: %s", err)
	}

	return grpc.Dial(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth),
	)
	/*
	   if err != nil {
	           return fmt.Errorf("did not connect: %s", err)
	   }
	   defer conn.Close()

	   c := pb.NewActionWorkflowClient(conn)

	   response, err := c.New(context.Background(), &pb.ActionRequest{Identifier: "foo"})
	   if err != nil {
	           return fmt.Errorf("Error when calling Get: %s", err)
	   }

	   log.Printf("Response from server: %s", response)
	*/
}
