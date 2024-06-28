/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// accessAndReadCmd represents the accessAndRead command
var accessAndReadCmd = &cobra.Command{
	Use:   "accessAndRead",
	Short: "Accesses tester repo, version number file, and returns contents",
	Long: `This command is responsible for accessing the VERSION_NUMBER file within the tester repo. 
	Returns the contents found within the file.`,

	Run: func(cmd *cobra.Command, args []string) {
		// hardcoded info: repo , file, owner

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accessAndReadCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accessAndReadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
