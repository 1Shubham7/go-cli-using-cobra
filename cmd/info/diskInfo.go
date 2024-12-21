/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ricochet2200/go-disk-usage/du"
)

// diskInfoCmd represents the diskInfo command
var diskInfoCmd = &cobra.Command{
	Use:   "diskInfo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskInfo called")

		usage := du.NewDiskUsage(".")
		fmt.Println("Available: ", usage.Available())
	},
}

func init() {
	InfoCmd.AddCommand(diskInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
