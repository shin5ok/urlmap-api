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

// createuserCmd represents the createuser command
var createuserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		notifyTo, _ := cmd.Flags().GetString("notify_to")
		host, _ := cmd.Flags().GetString("host")

		Client := CreateClient(host)

		u := &pb.User{User: user, NotifyTo: notifyTo}
		if res, err := Client.SetUser(context.TODO(), u); err != nil {
			log.Printf("%+v\n", err)
		} else {
			j, _ := json.MarshalIndent(res, "", " ")
			fmt.Println(string(j))
		}
	},
}

func init() {
	rootCmd.AddCommand(createuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createuserCmd.Flags().String("user", "", "")
	createuserCmd.Flags().String("notify_to", "", "")
}
