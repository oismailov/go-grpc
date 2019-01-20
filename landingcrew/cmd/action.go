package cmd

import (
	"context"
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/grpc_client"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"os"
)

func getCmdAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action",
		Short: "Handles actions.",
		Long:  "",
	}

	conn, ctx, cancelFunc := grpc_client.GetConnectionSetting()

	client := pb.NewActionWorkflowClient(conn)

	cmd.AddCommand(getCmdActionList(client, ctx, cancelFunc, conn))
	cmd.AddCommand(geCmdActionGet())

	return cmd
}

func getCmdActionList(client pb.ActionWorkflowClient, ctx context.Context, cancelFunc context.CancelFunc, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all actions.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			defer conn.Close()
			defer cancelFunc()

			stream, err := client.List(ctx, &empty.Empty{})

			if err != nil {
				log.Fatalf("Could not read from stream: %s", err)
			}

			waitc := make(chan struct{})

			go func() {
				for {
					r, err := stream.Recv()

					if err != nil {
						os.Exit(1)
					}

					fmt.Printf("%s", lib.ConvertStructToJson(&lib.Output{Error: r.Message}))
				}
			}()
			<-waitc
			stream.CloseSend()
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
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of action that will be shown.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdAction())
}
