package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"

	pb "github.com/abhiyerra/landingcrew-cli/landingcrew/workflow"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "landingcrew",
	Short: "Here will be short description",
	Long:  "Here will be long description",
}

var (
	ActionWorkflowClient  pb.ActionWorkflowClient
	CodeWorkflowClient    pb.CodeWorkflowClient
	ContentWorkflowClient pb.ContentWorkflowClient
	DataWorkflowClient    pb.DataWorkflowClient
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(conn *grpc.ClientConn) {
	defer conn.Close()

	ActionWorkflowClient = pb.NewActionWorkflowClient(conn)
	CodeWorkflowClient = pb.NewCodeWorkflowClient(conn)
	ContentWorkflowClient = pb.NewContentWorkflowClient(conn)
	DataWorkflowClient = pb.NewDataWorkflowClient(conn)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
