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

func getCmdCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "Handles coding tasks.",
		Long:  "",
	}

	cmd.AddCommand(getCmdCodeNew())
	cmd.AddCommand(getCmdCodeInit())
	cmd.AddCommand(geCmdCodeGet())
	cmd.AddCommand(getCmdCodeList())
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

		},
	}

	cmd.Flags().StringVar(&githubAuthToken, "github-auth-token", "", "github auth token.")
	err := cmd.MarkFlagRequired("github-auth-token")
	if err != nil {
		log.Fatalf("Could not mark flag `github-auth-token` as required: %s", err)
	}

	cmd.Flags().StringVar(&githubRepo, "github-repo", "", "github repo.")
	err = cmd.MarkFlagRequired("github-repo")
	if err != nil {
		log.Fatalf("Could not mark flag `github-repo` as required: %s", err)
	}

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	err = cmd.MarkFlagRequired("code-type")
	if err != nil {
		log.Fatalf("Could not mark flag `code-type` as required: %s", err)
	}

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

		},
	}

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	err := cmd.MarkFlagRequired("code-type")
	if err != nil {
		log.Fatalf("Could not mark flag `code-type` as required: %s", err)
	}

	cmd.Flags().StringVar(&name, "name", "", "name.")
	err = cmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Could not mark flag `name` as required: %s", err)
	}

	return cmd
}

func geCmdCodeGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := codeWorkflowClient.Get(context.Background(), &pb.CodeRequest{Id: id})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that shown")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Could not mark flag `id` as required: %s", err)
	}

	return cmd
}

func getCmdCodeApprove() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "approve [options]",
		Short: "Approve coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that will be approved")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Could not mark flag `id` as required: %s", err)
	}

	return cmd
}

func getCmdCodeList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all coding tasks.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			stream, err := codeWorkflowClient.List(context.Background(), &empty.Empty{})

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

func init() {
	rootCmd.AddCommand(getCmdCode())
}
