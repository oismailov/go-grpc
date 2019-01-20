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

func getCmdCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "Handles coding tasks.",
		Long:  "",
	}

	conn, ctx, cancelFunc := grpc_client.GetConnectionSetting()

	client := pb.NewCodeWorkflowClient(conn)

	cmd.AddCommand(getCmdCodeNew())
	cmd.AddCommand(getCmdCodeInit())
	cmd.AddCommand(geCmdCodeGet())
	cmd.AddCommand(getCmdCodeList(client, ctx, cancelFunc, conn))
	cmd.AddCommand(getCmdCodeApprove())

	return cmd
}

func getCmdCodeNew() *cobra.Command {
	var githubAuthToken string
	var githubRepo string
	var codeType string

	cmd := &cobra.Command{
		Use:   "new [options]",
		Short: "Create new coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&githubAuthToken, "github-auth-token", "", "github auth token.")
	cmd.MarkFlagRequired("github-auth-token")

	cmd.Flags().StringVar(&githubRepo, "github-repo", "", "github repo.")
	cmd.MarkFlagRequired("github-repo")

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	cmd.MarkFlagRequired("code-type")

	return cmd
}

func getCmdCodeInit() *cobra.Command {
	var codeType string
	var name string

	cmd := &cobra.Command{
		Use:   "init [options]",
		Short: "Init coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	cmd.MarkFlagRequired("code-type")

	cmd.Flags().StringVar(&name, "name", "", "name.")
	cmd.MarkFlagRequired("name")

	return cmd
}

func geCmdCodeGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that shown")
	cmd.MarkFlagRequired("id")

	return cmd
}

func getCmdCodeApprove() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "approve [options]",
		Short: "Approve coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that will be approved")
	cmd.MarkFlagRequired("id")

	return cmd
}

func getCmdCodeList(client pb.CodeWorkflowClient, ctx context.Context, cancelFunc context.CancelFunc, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all coding tasks.",
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

func init() {
	rootCmd.AddCommand(getCmdCode())
}
