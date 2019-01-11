package cmd

import (
	"fmt"

	"github.com/landingcrew-cli/landingcrew/lib"
	"github.com/landingcrew-cli/landingcrew/utils"
	"github.com/spf13/cobra"
)

type Code struct {
	ID         string
	Identifier string `json:",omitempty"`
	Status     lib.Status
	Fields     map[string]interface{}
	Bid        int32  `json:",omitempty"`
	Error      string `json:",omitempty"`
}

//define struct to store values from CLI input (flags)
type Args struct {
	GithubAuthToken string
	GithubRepo      string
	GithubCodeType  string
	CodeId          string
	Name            string
}

//define command functions
func getCmdCode() *cobra.Command {
	var cmdCode = &cobra.Command{
		Use:   "code",
		Short: "Handles coding tasks.",
		Long:  "",
	}

	cmdCode.AddCommand(getCmdCodeNew())
	cmdCode.AddCommand(getCmdCodeInit())
	cmdCode.AddCommand(geCmdCodeGet())
	cmdCode.AddCommand(getCmdCodeList())
	cmdCode.AddCommand(getCmdCodeApprove())

	return cmdCode
}

func getCmdCodeNew() *cobra.Command {
	cmdArgs := new(Args)

	var cmdCodeNew = &cobra.Command{
		Use:   "new [options]",
		Short: "Create new coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.ConvertStructToJson(&Code{}))
		},
	}

	cmdCodeNew.Flags().StringVar(&cmdArgs.GithubAuthToken, "github-auth-token", "", "github auth token.")
	cmdCodeNew.MarkFlagRequired("github-auth-token")

	cmdCodeNew.Flags().StringVar(&cmdArgs.GithubRepo, "github-repo", "", "github repo.")
	cmdCodeNew.MarkFlagRequired("github-repo")

	cmdCodeNew.Flags().StringVar(&cmdArgs.GithubCodeType, "code-type", "", "github code type.")
	cmdCodeNew.MarkFlagRequired("code-type")

	return cmdCodeNew
}

func getCmdCodeInit() *cobra.Command {
	cmdArgs := new(Args)

	var cmdCodeInit = &cobra.Command{
		Use:   "init [options]",
		Short: "Init coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.ConvertStructToJson(&Code{}))
		},
	}

	cmdCodeInit.Flags().StringVar(&cmdArgs.GithubCodeType, "code-type", "", "github code type.")
	cmdCodeInit.MarkFlagRequired("code-type")

	cmdCodeInit.Flags().StringVar(&cmdArgs.Name, "name", "", "name.")
	cmdCodeInit.MarkFlagRequired("name")

	return cmdCodeInit
}

func geCmdCodeGet() *cobra.Command {
	cmdArgs := new(Args)

	var cmdCodeGet = &cobra.Command{
		Use:   "get [options]",
		Short: "Show single coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.ConvertStructToJson(&Code{}))
		},
	}

	cmdCodeGet.Flags().StringVar(&cmdArgs.CodeId, "id", "", "Id of code that will be approved")
	cmdCodeGet.MarkFlagRequired("id")

	return cmdCodeGet
}

func getCmdCodeApprove() *cobra.Command {
	cmdArgs := new(Args)

	var cmdCodeApprove = &cobra.Command{
		Use:   "approve [options]",
		Short: "Approve coding task.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.ConvertStructToJson(&Code{}))
		},
	}

	cmdCodeApprove.Flags().StringVar(&cmdArgs.CodeId, "id", "", "Id of code that will be approved")
	cmdCodeApprove.MarkFlagRequired("id")

	return cmdCodeApprove
}

func getCmdCodeList() *cobra.Command {
	var cmdCodeList = &cobra.Command{
		Use:   "list",
		Short: "Show all coding tasks.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.ConvertStructToJson(&Code{}))
		},
	}

	return cmdCodeList
}

func init() {
	rootCmd.AddCommand(getCmdCode())
}
