package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdCode = &cobra.Command{
		Use:   "code [options]",
		Short: "Create new coding task.",
		Long: `Create a new coding task/

This creates a new Github Repo with the content of a director that you want.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdCodeNew = &cobra.Command{
		Use:   "new [options]",
		Short: "Create new coding task.",
		Long: `Create a new coding task/

This creates a new Github Repo with the content of a director that you want.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	cmdCode.AddCommand(cmdCodeNew)

	var cmdContent = &cobra.Command{
		Use:   "content [options]",
		Short: "Create new content task.",
		Long: `Create a new coding task/

This creates a new Github Repo with the content of a director that you want.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdData = &cobra.Command{
		Use:   "data [options]",
		Short: "Create new data task.",
		Long: `Create a new coding task/

This creates a new Github Repo with the content of a director that you want.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdAction = &cobra.Command{
		Use:   "action [options]",
		Short: "Create new data task.",
		Long: `Create a new coding task/

This creates a new Github Repo with the content of a director that you want.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var rootCmd = &cobra.Command{Use: "landingcrew"}
	rootCmd.AddCommand(cmdCode, cmdContent, cmdData, cmdAction)
	rootCmd.Execute()
}
