package cmd

import (
	"context"
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

func getCmdAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action",
		Short: "Handles actions.",
		Long:  "",
	}

	cmd.AddCommand(getCmdActionList())
	cmd.AddCommand(geCmdActionGet())

	return cmd
}

func getCmdActionList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all actions.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			stream, err := actionWorkflowClient.List(context.Background(), &empty.Empty{})

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

func geCmdActionGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single action.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := actionWorkflowClient.Get(context.Background(), &pb.ActionRequest{Id: id})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of action that will be shown.")
	err := cmd.MarkFlagRequired("id")

	if err != nil {
		log.Fatalf("Could not mark flag `id` as required: %s", err)
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdAction())
}
