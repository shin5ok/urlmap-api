/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/shin5ok/urlmap-api/pb"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		user, _ := cmd.Flags().GetString("user")
		Client = CreateClient(host)

		u := &pb.User{User: user}
		if res, err := Client.GetInfoByUser(context.TODO(), u); err != nil {
			log.Printf("error:%#v \n", err)
		} else {
			j, _ := json.MarshalIndent(res, "", " ")
			fmt.Println(string(j))
		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().String("user", "", "")
}
