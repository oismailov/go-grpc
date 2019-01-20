package cmd

import (
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/grpc_client"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"os"

	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	"github.com/spf13/cobra"

	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"

	"context"
)

func getCmdContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content",
		Short: "Handles content tasks.",
		Long:  "",
	}

	conn, ctx, cancelFunc := grpc_client.GetConnectionSetting()

	client := pb.NewContentWorkflowClient(conn)

	cmd.AddCommand(getCmdContentList(client, ctx, cancelFunc, conn))
	cmd.AddCommand(geCmdContentGet())
	cmd.AddCommand(getCmdContentNew())
	cmd.AddCommand(getCmdContentInit())
	cmd.AddCommand(getCmdContentTypeList())

	return cmd
}

func getCmdContentList(client pb.ContentWorkflowClient, ctx context.Context, cancelFunc context.CancelFunc, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all content tasks.",
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

func geCmdContentGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single content task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of content that will be shown.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func getCmdContentNew() *cobra.Command {
	var contentType string
	var file string

	cmd := &cobra.Command{
		Use:   "new [options]",
		Short: "Create new content task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&contentType, "type", "", "content type.")
	cmd.MarkFlagRequired("type")

	cmd.Flags().StringVar(&file, "file", "", "file name.")
	cmd.MarkFlagRequired("file")

	return cmd
}

func getCmdContentInit() *cobra.Command {
	var contentType string
	var name string

	cmd := &cobra.Command{
		Use:   "init [options]",
		Short: "Init content task",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&contentType, "content-type", "", "content type.")
	cmd.MarkFlagRequired("content-type")

	cmd.Flags().StringVar(&name, "name", "", "content name")
	cmd.MarkFlagRequired("name")

	return cmd
}

func getCmdContentTypeList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "type-list [options]",
		Short: "List content type",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdContent())
}
