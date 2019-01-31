package cmd

import (
	"context"
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/lib"
	pb "github.com/abhiyerra/landingcrew-cli/workflow"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

func getCmdApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Handles app request.",
		Long:  "",
	}

	cmd.AddCommand(getCmdAppList())
	cmd.AddCommand(getCmdAppInvoke())

	return cmd
}

func getCmdAppList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List applications.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			stream, err := appWorkflowClient.List(context.Background(), &empty.Empty{})

			if err != nil {
				log.Fatalf("Could not read from stream: %s", err)
			}

			for {
				response, err := stream.Recv()

				if err != nil {
					if err == io.EOF {
						break
					}
					os.Exit(1)
				}

				fmt.Printf("%v", lib.ConvertStructToJson(response))
			}
		},
	}

	return cmd
}

func getCmdAppInvoke() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Invoke application.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := appWorkflowClient.Invoke(context.Background(), &pb.AppRequest{})

			if err != nil {
				log.Fatalf("Could not invoke application: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdApp())
}
