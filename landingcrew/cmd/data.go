package cmd

import (
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/grpc_client"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"

	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	"github.com/spf13/cobra"

	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"

	"context"
	"os"
)

func getCmdData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "Handles actions.",
		Long:  "",
	}

	conn, ctx, cancelFunc := grpc_client.GetConnectionSetting()

	client := pb.NewDataWorkflowClient(conn)

	cmd.AddCommand(getCmdDataList(client, ctx, cancelFunc, conn))
	cmd.AddCommand(geCmdDataGet())
	cmd.AddCommand(getCmdDataNew())

	return cmd
}

func getCmdDataList(client pb.DataWorkflowClient, ctx context.Context, cancelFunc context.CancelFunc, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all data.",
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

func geCmdDataGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single data.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
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
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdData())
}
