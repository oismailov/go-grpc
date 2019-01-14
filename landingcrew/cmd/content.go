package cmd

import (
	"fmt"

	"github.com/landingcrew-cli/landingcrew/lib"
	"github.com/landingcrew-cli/landingcrew/utils"
	"github.com/spf13/cobra"
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
			fmt.Println(utils.ConvertStructToJson(&lib.Output{}))
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
			fmt.Println(utils.ConvertStructToJson(&lib.Output{}))
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
			fmt.Println(utils.ConvertStructToJson(&lib.Output{}))
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
			fmt.Println(utils.ConvertStructToJson(&lib.Output{}))
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
			fmt.Println(utils.ConvertStructToJson(&lib.Output{}))
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(getCmdContent())
}
