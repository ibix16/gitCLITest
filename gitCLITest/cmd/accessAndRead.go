/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// accessAndReadCmd represents the accessAndRead command
var accessAndReadCmd = &cobra.Command{
	Use:   "accessAndRead",
	Short: "Accesses tester repo, version number file, and returns contents",
	Long: `This command is responsible for accessing the VERSION_NUMBER file within the tester repo. 
	Returns the contents found within the file.`,

	Run: func(cmd *cobra.Command, args []string) {
		// hardcoded info: repo , file, owner
		owner := "ibix16"
		repo := "tester"
		filePath := "VERSION_NUMBER"

		err := accessAndRead(owner, repo, filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	},
}

func accessAndRead(repoOwner, repoName, filePath string) error {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	//retrieve file contents
	fileContent, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, filePath, nil)
	if err != nil {
		return err
	}

	// print contents
	content, err := fileContent.GetContent()
	if err != nil {
		return err
	}
	fmt.Println(content)

	return nil

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
