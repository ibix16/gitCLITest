/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// bumpCmd represents the bump command
var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "This command is responsible for accessing the VERSION_NUMBER file and incrementing the version number by 1. Returns the updated version number",
	Long: `This command is responsible for accessing the 
	VERSION_NUMBER file and incrementing the version number by 1. 
	Returns the updated version number.`,

	Run: func(cmd *cobra.Command, args []string) {
		owner := "ibix16"
		repo := "tester"
		filePath := "VERSION_NUMBER"

		err := bump(owner, repo, filePath)

		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// cmd to bump version number
func bump(owner, repo, filePath string) error {
	// Get the current file content
	content, err := getFileContents(owner, repo, filePath)
	if err != nil {
		return err
	}

	// Increment the version number
	updatedContent := incrementVersionNumber(content)

	// Print or process the updated content
	fmt.Println(updatedContent)

	return nil

}

// gets file content via api
func getFileContents(repoOwner, repoName, filePath string) (string, error) {
	//client := GitHubConfig.NewClient()

	// create context for api call
	ctx := context.Background()

	//retrieve file contents
	fileContent, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, filePath, nil)
	if err != nil {
		fmt.Print(err)
	}

	// print contents
	content, err := fileContent.GetContent()
	if err != nil {
		fmt.Print(err)
	}

	return content, nil
}

// parses version number
// isolates patch version by searching for last digit seperated by .
// increments isolated variable
// rejoines string + returns string
func incrementVersionNumber(content string) string {
	// Parse the version number from the content
	//parts := strings.Split(content, ".")
	//versionPart, _ := strconv.Atoi(parts[len(parts)-1])

	// Increment the version number
	//versionPart = versionPart + 2

	// Reconstruct the updated content
	//parts[len(parts)-1] = strconv.Itoa(versionPart)
	//updatedContent := strings.Join(parts, ".")

	//testVar := strconv.Itoa(versionPart)
	return content
	//updatedContent
}

func init() {
	rootCmd.AddCommand(bumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
