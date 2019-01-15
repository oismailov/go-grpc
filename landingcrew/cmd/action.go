package cmd

import (
	"fmt"

	"github.com/landingcrew-cli/landingcrew/lib"
	"github.com/spf13/cobra"
)

func getCmdAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action",
		Short: "Handles actions.",
		Long:  "",
	}

	cmd.AddCommand(getCmdActionList())
	cmd.AddCommand(geCmdActionGet())

	return cmd
}

func getCmdActionList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all actions.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lib.ConvertStructToJson(&lib.Output{}))
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
