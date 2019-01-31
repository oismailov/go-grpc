package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"

	pb "github.com/abhiyerra/landingcrew-cli/workflow"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "landingcrew",
	Short: "Here will be short description",
	Long:  "Here will be long description",
}

var (
	actionWorkflowClient  pb.ActionWorkflowClient
	codeWorkflowClient    pb.CodeWorkflowClient
	contentWorkflowClient pb.ContentWorkflowClient
	dataWorkflowClient    pb.DataWorkflowClient
	appWorkflowClient     pb.AppWorkflowClient
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(conn *grpc.ClientConn) {
	defer conn.Close()

	actionWorkflowClient = pb.NewActionWorkflowClient(conn)
	codeWorkflowClient = pb.NewCodeWorkflowClient(conn)
	contentWorkflowClient = pb.NewContentWorkflowClient(conn)
	dataWorkflowClient = pb.NewDataWorkflowClient(conn)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
