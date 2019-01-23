package cmd

import (
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"

	"github.com/spf13/cobra"

	"context"
	"os"
)

func getCmdData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "Handles actions.",
		Long:  "",
	}

	cmd.AddCommand(getCmdDataList())
	cmd.AddCommand(geCmdDataGet())
	cmd.AddCommand(getCmdDataNew())

	return cmd
}

func getCmdDataList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all data.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			stream, err := DataWorkflowClient.List(context.Background(), &empty.Empty{})

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

func geCmdDataGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single data.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of data that will be shown.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func getCmdDataNew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [options]",
		Short: "Create new data.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdData())
}
