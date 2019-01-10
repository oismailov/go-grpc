package main

import (
	"github.com/landingcrew-cli/landingcrew/cmd"
)

func main() {
	// 	var cmdContent = &cobra.Command{
	// 		Use:   "content [options]",
	// 		Short: "Create new content task.",
	// 		Long: `Create a new coding task/
	//
	// This creates a new Github Repo with the content of a director that you want.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			fmt.Println("Print: " + strings.Join(args, " "))
	// 		},
	// 	}
	//
	// 	var cmdData = &cobra.Command{
	// 		Use:   "data [options]",
	// 		Short: "Create new data task.",
	// 		Long: `Create a new coding task/
	//
	// This creates a new Github Repo with the content of a director that you want.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			fmt.Println("Print: " + strings.Join(args, " "))
	// 		},
	// 	}
	//
	// 	var cmdAction = &cobra.Command{
	// 		Use:   "action [options]",
	// 		Short: "Create new data task.",
	// 		Long: `Create a new coding task/
	//
	// This creates a new Github Repo with the content of a director that you want.`,
	// 		Args: cobra.MinimumNArgs(1),
	// 		Run: func(cmd *cobra.Command, args []string) {
	// 			fmt.Println("Print: " + strings.Join(args, " "))
	// 		},
	// 	}
	//

	cmd.Execute()
}
