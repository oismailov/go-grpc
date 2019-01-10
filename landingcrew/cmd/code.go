package cmd

import (
	"fmt"

	"github.com/landingcrew-cli/landingcrew/utils"
	"github.com/spf13/cobra"
)

type Status struct{}

type Code struct {
	ID         string
	Identifier string `json:",omitempty"`
	Status     Status
	Fields     map[string]interface{}
	Bid        int32  `json:",omitempty"`
	Error      string `json:",omitempty"`
}

//define variables to store values from CLI input (flags)
var githubAuthToken string
var githubRepo string
var githubCodeType string
var codeId string
var name string

//define commands
var cmdCode = &cobra.Command{
	Use:   "code",
	Short: "Handles coding tasks.",
	Long:  "",
}

var cmdCodeNew = &cobra.Command{
	Use:   "new [options]",
	Short: "Create new coding task.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//create new one
		fmt.Println(utils.ConvertStructToJson(&Code{}))
	},
}

var cmdCodeInit = &cobra.Command{
	Use:   "init [options]",
	Short: "Init coding task.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.ConvertStructToJson(&Code{}))
	},
}

var cmdCodeGet = &cobra.Command{
	Use:   "get [options]",
	Short: "Show single coding task.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//find one by id
		fmt.Println(utils.ConvertStructToJson(&Code{ID: codeId}))
	},
}

var cmdCodeApprove = &cobra.Command{
	Use:   "approve [options]",
	Short: "Approve coding task.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.ConvertStructToJson(&Code{ID: codeId}))
	},
}

var cmdCodeList = &cobra.Command{
	Use:   "list",
	Short: "Show all coding tasks.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//find all
		fmt.Println(utils.ConvertStructToJson(&Code{}))
	},
}

func init() {
	//define flags for `code new` command
	cmdCodeNew.Flags().StringVar(&githubAuthToken, "github-auth-token", "", "github auth token.")
	cmdCodeNew.MarkFlagRequired("github-auth-token")

	cmdCodeNew.Flags().StringVar(&githubRepo, "github-repo", "", "github repo.")
	cmdCodeNew.MarkFlagRequired("github-repo")

	cmdCodeNew.Flags().StringVar(&githubCodeType, "code-type", "", "github code type.")
	cmdCodeNew.MarkFlagRequired("code-type")

	//define flags for `code approve` command
	cmdCodeApprove.Flags().StringVar(&codeId, "id", "", "Id of code that will be approved")
	cmdCodeApprove.MarkFlagRequired("id")

	//define flags for `code get` command
	cmdCodeGet.Flags().StringVar(&codeId, "id", "", "Id of code that will be approved")
	cmdCodeGet.MarkFlagRequired("id")

	//define flags for `code init` command
	cmdCodeInit.Flags().StringVar(&githubCodeType, "code-type", "", "github code type.")
	cmdCodeInit.MarkFlagRequired("code-type")

	cmdCodeInit.Flags().StringVar(&name, "name", "", "name.")
	cmdCodeInit.MarkFlagRequired("name")

	//add commands to rootCmd
	cmdCode.AddCommand(cmdCodeNew)
	cmdCode.AddCommand(cmdCodeInit)
	cmdCode.AddCommand(cmdCodeGet)
	cmdCode.AddCommand(cmdCodeList)
	cmdCode.AddCommand(cmdCodeApprove)
	rootCmd.AddCommand(cmdCode)
}
