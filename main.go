/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ibix16/gitCLITest/config"
	"github.com/ibix16/gitCLITest/gitCLITest/cmd"
)

func main() {

	err := config.LoadGithubConfig()
	if err != nil {
		// Handle the error
		return
	}

	cmd.Execute()
}
