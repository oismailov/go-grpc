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
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const CODE_TEMPLATE_PATH = "https://assets.landingcrew.com/templates/{CODE_TYPE}.zip"

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
	cmd.AddCommand(getCmdCodeTypeList())
	cmd.AddCommand(getCmdCodeReject())

	return cmd
}

func getCmdCodeTypeList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "type-list [options]",
		Short: "Output list of types.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("%v", lib.ConvertStructToJson(pb.CodeType_value))
		},
	}

	return cmd
}

func getCmdCodeNew() *cobra.Command {
	var (
		githubAuthToken string
		githubRepo      string
		codeType        string
	)

	cmd := &cobra.Command{
		Use:   "new [options]",
		Short: "Create new coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			codeTypeValue, ok := pb.CodeType_value[codeType]

			if !ok {
				log.Fatalf("invalid enum value: %s", codeType)
			}

			response, err := codeWorkflowClient.New(context.Background(), &pb.CodeRequest{
				Github: &pb.Github{
					AuthToken: githubAuthToken,
					FullRepo:  githubRepo,
				},
				Type: pb.CodeType(codeTypeValue),
			})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&githubAuthToken, "github-auth-token", "", "github auth token.")
	if err := cmd.MarkFlagRequired("github-auth-token"); err != nil {
		log.Fatalf("Could not mark flag `github-auth-token` as required: %s", err)
	}

	cmd.Flags().StringVar(&githubRepo, "github-repo", "", "github repo.")
	if err := cmd.MarkFlagRequired("github-repo"); err != nil {
		log.Fatalf("Could not mark flag `github-repo` as required: %s", err)
	}

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	if err := cmd.MarkFlagRequired("code-type"); err != nil {
		log.Fatalf("Could not mark flag `code-type` as required: %s", err)
	}

	return cmd
}

func getCmdCodeInit() *cobra.Command {
	var (
		codeType            string
		name                string
		path                string
		tmpDownloadFilePath string
	)

	cmd := &cobra.Command{
		Use:   "init [options]",
		Short: "Init coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			/*
				Create tmp zip file `tmpDownloadFilePath`.
				Download archive from `https://assets.landingcrew.com/templates/{CODE_TYPE}/.zip` to that file.
				Unzip `tmpDownloadFilePath` to `path`
				Replace  {{Name .}} with name for each file.
			*/
			tmpDownloadFilePath = strconv.Itoa(time.Now().Nanosecond()*rand.Int()) + ".zip"

			downloadFilePath := strings.Replace(CODE_TEMPLATE_PATH, "{CODE_TYPE}", codeType, -1)

			if err := lib.DownloadFile(tmpDownloadFilePath, downloadFilePath); err != nil {
				log.Fatalf("Could not download file %s: %s", downloadFilePath, err)
			}

			files, err := lib.Unzip(tmpDownloadFilePath, path)
			if err != nil {
				log.Fatalf("Could not unzip archive: %s", err)
			}

			if err := os.Remove(tmpDownloadFilePath); err != nil {
				log.Fatalf("Could not remove tmp file %s : %s", tmpDownloadFilePath, err)
			}

			lib.FindReplaceText(files, name)
		},
	}

	cmd.Flags().StringVar(&codeType, "code-type", "", "github code type.")
	if err := cmd.MarkFlagRequired("code-type"); err != nil {
		log.Fatalf("Could not mark flag `code-type` as required: %s", err)
	}

	cmd.Flags().StringVar(&name, "name", "", "name.")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Could not mark flag `name` as required: %s", err)
	}

	cmd.Flags().StringVar(&path, "path", "", "path.")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Could not mark flag `path` as required: %s", err)
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
			response, err := codeWorkflowClient.Get(context.Background(), &pb.CodeRequest{Id: lib.ConvertStringToInt64(id)})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that shown")

	if err := cmd.MarkFlagRequired("id"); err != nil {
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
			response, err := codeWorkflowClient.Decide(context.Background(), &pb.CodeDecision{
				Id:      lib.ConvertStringToInt64(id),
				Decison: pb.CodeDecisionType_APPROVE,
			})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that will be approved")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("Could not mark flag `id` as required: %s", err)
	}

	return cmd
}

func getCmdCodeReject() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "reject [options]",
		Short: "Reject coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := codeWorkflowClient.Decide(context.Background(), &pb.CodeDecision{
				Id:      lib.ConvertStringToInt64(id),
				Decison: pb.CodeDecisionType_REJECT,
			})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of code that will be rejected")
	if err := cmd.MarkFlagRequired("id"); err != nil {
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
