/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	pb "github.com/shin5ok/urlmap-api/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var Client pb.RedirectionClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A brief description of your application",
	Long:  "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		Client = CreateClient(host)
	},
}

func CreateClient(host string) pb.RedirectionClient {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	// defer conn.Close()
	Client = pb.NewRedirectionClient(conn)
	return Client

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("host", "localhost:8080", "")
}

func initConfig() {
}
