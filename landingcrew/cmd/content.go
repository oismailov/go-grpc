package cmd

import (
	"fmt"
	"github.com/abhiyerra/landingcrew-cli/landingcrew/lib"
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

		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdContent())
}
