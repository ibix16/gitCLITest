/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

var (
	targetBranchT  = "version-update"
	commitMessageT = "Bump version number to %s"
	ownerT         = "ibix16"
	repoT          = "tester"
	filePathT      = "/VERSION_NUMBER"
)

// bumpCmd represents the bump command
var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "This command is responsible for accessing the VERSION_NUMBER file and incrementing the version number by 1.",
	Long: `This command is responsible for accessing the 
	VERSION_NUMBER file and incrementing the version number by 1. 
	Commits changes and raises PR.`,

	Run: func(cmd *cobra.Command, args []string) {

		err := bump(ownerT, repoT, filePathT)
		if err != nil {
		 log.Fatal("err", err)
		}

	},
}

func bump(owner, repo, filePath string) error {
	// create client
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(accessToken)

	// get file contents
	fileContent, err := getFileContents(owner, repo, filePath)
	if err != nil {
		fmt.Errorf("invalid version number format: %s", fileContent)
	}
	// returns string "v0.5.6"

	// split retured string v0.0.0
	parts := strings.Split(fileContent, ".")
	if len(parts) != 3 {
		fmt.Errorf("invalid version number format: %s", fileContent)
	}
	// parts[0] holds the major value : v0
	// parts[1] holds the minor value : 5
	// parts[2] holds the patch value : 6

	// increment patch value
	numstr := parts[2]
	trimmedValue := strings.TrimSpace(numstr)
	convertedValue, err := strconv.Atoi(trimmedValue)
	if err != nil {
		fmt.Print(err)
	}
	convertedValue++

	// concatenate string back together correctly bumps version from "v0.5.6" to "v0.5.7"
	updatedString := fmt.Sprintf("%s.%s.%d", parts[0], parts[1], convertedValue)



	// get latest commit sha
	ref, _, err := client.Git.GetRef(ctx, ownerT, repoT, "heads/"+targetBranchT)
	if err != nil {
		return fmt.Errorf("error getting ref %s", err)
	}
	latestCommitSha := ref.Object.GetSHA()

	entries := []*github.TreeEntry{}
	entries = append(entries, &github.TreeEntry{Path: github.String(strings.TrimPrefix(filePath, "/")), Type: github.String("blob"), Content: github.String(string(updatedString)), Mode: github.String("100644")})
	tree, _, err := client.Git.CreateTree(ctx,ownerT, repoT, *ref.Object.SHA, entries)
	if err != nil {
		return fmt.Errorf("creating tree %s", err)
	}


	//validate tree sha
	newTreeSHA := tree.GetSHA()


	// create new commit
	author := &github.CommitAuthor{
		Name:  github.String("ibix16"),
		Email: github.String("ibixrivera16@gmail.com"),
	}
	commit := &github.Commit{
		Message: github.String(fmt.Sprintf(commitMessageT, updatedString)),
		Tree:    &github.Tree{SHA: github.String(newTreeSHA)},
		Author:  author,
		Parents: []*github.Commit{{SHA: github.String(latestCommitSha)}},
	}
	commitOP := &github.CreateCommitOptions{}
	newCommit, _, err := client.Git.CreateCommit(ctx, ownerT, repoT, commit, commitOP)
	if err != nil {
		return fmt.Errorf("creating commit %s", err)
	}
	newCommitSHA := newCommit.GetSHA()


	// update branch reference
	ref.Object.SHA = github.String(newCommitSHA)


	_, _, err = client.Git.UpdateRef(ctx, ownerT, repoT, ref, false)
	if err != nil {
		return fmt.Errorf("error updating ref %s", err)
	}

	
	// create pull request
    base := "main"
    head := fmt.Sprintf("%s:%s", ownerT, targetBranchT)
    title := fmt.Sprintf("Version bump to %s", updatedString)
    body := "This pull request bumps the version number."

    newPR := &github.NewPullRequest{
        Title: &title,
        Head:  &head,
        Base:  &base,
        Body:  &body,
    }

	pr, _, err := client.PullRequests.Create(ctx, owner, repo, newPR)
    if err != nil {
        return err
    }


	log.Printf("Pull request created: %s\n", pr.GetHTMLURL())
	return err

}

// gets file content via api
func getFileContents(repoOwner, repoName, filePath string) (string, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(accessToken)

	//retrieve file contents
	fileContent, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, filePath, nil)
	if err != nil {
		fmt.Print(err)
	}

	content, err := fileContent.GetContent()
	if err != nil {
		fmt.Print(err)
	}

	return content, nil
}

