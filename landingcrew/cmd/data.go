package cmd

import (
	"fmt"

	"github.com/abhiyerra/landingcrew-cli/lib"
	"github.com/spf13/cobra"
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
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
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
