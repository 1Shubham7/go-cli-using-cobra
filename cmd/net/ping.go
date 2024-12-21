/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "this pings a remote URL and returns the response",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")
		// Logic comes here
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The URL to ping")
	NetCmd.AddCommand(pingCmd)

	// to set a flag as required
	err := NetCmd.MarkFlagRequired("url")
	if err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
