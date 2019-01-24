package cmd

import (
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"context"
)

func getCmdContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content",
		Short: "Handles content tasks.",
		Long:  "",
	}

	cmd.AddCommand(getCmdContentList())
	cmd.AddCommand(geCmdContentGet())
	cmd.AddCommand(getCmdContentNew())
	cmd.AddCommand(getCmdContentInit())
	cmd.AddCommand(getCmdContentTypeList())

	return cmd
}

func getCmdContentList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all content tasks.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			stream, err := contentWorkflowClient.List(context.Background(), &empty.Empty{})
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

func geCmdContentGet() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get [options]",
		Short: "Show single content task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := contentWorkflowClient.Get(context.Background(), &pb.ContentRequest{Id: id})

			if err != nil {
				log.Fatalf("Could not get response from server: %s", err)
			}

			fmt.Printf("%v", lib.ConvertStructToJson(response))
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Id of content that will be shown.")

	if err := cmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("Could not mark flag `id` as required: %s", err)
	}

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

		},
	}

	cmd.Flags().StringVar(&contentType, "type", "", "content type.")
	if err := cmd.MarkFlagRequired("type"); err != nil {
		log.Fatalf("Could not mark flag `type` as required: %s", err)
	}

	cmd.Flags().StringVar(&file, "file", "", "file name.")
	if err := cmd.MarkFlagRequired("file"); err != nil {
		log.Fatalf("Could not mark flag `file` as required: %s", err)
	}

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

		},
	}

	cmd.Flags().StringVar(&contentType, "content-type", "", "content type.")
	if err := cmd.MarkFlagRequired("content-type"); err != nil {
		log.Fatalf("Could not mark flag `content-type` as required: %s", err)
	}

	cmd.Flags().StringVar(&name, "name", "", "content name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Could not mark flag `name` as required: %s", err)
	}

	return cmd
}

func getCmdContentTypeList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "type-list [options]",
		Short: "List content type",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%v", lib.ConvertStructToJson(pb.ContentType_value))
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdContent())
}
