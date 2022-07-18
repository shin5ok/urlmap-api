/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	pb "github.com/shin5ok/urlmap-api/pb"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		user, _ := cmd.Flags().GetString("user")
		org, _ := cmd.Flags().GetString("org")
		Client = CreateClient(host)
		randPath := strings.Split(uuid.New().String(), "-")[0]

		data := &pb.RedirectData{}
		data.Redirect = &pb.RedirectInfo{
			User:         user,
			Org:          org,
			RedirectPath: randPath,
			Comment:      "sample test",
			Active:       1,
		}
		if res, err := Client.SetInfo(context.TODO(), data); err != nil {
			log.Println(data.Redirect)
			log.Printf("error::%#v \n", err)
		} else {
			log.Printf(randPath)
			log.Printf("result:%#v \n", res)
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setCmd.Flags().String("path", "", "")
	setCmd.Flags().String("org", "", "")
	setCmd.Flags().String("user", "", "")
}
