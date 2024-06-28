package config

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type GithubConfig struct {
	AccessToken string
	// Add any other GitHub-related configurations here

}

var githubConfig *GithubConfig

func LoadGithubConfig() error {
	viper.SetConfigName("github")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	githubConfig = &GithubConfig{
		AccessToken: viper.GetString("GITHUB_ACCESS_TOKEN"),
		// Load other GitHub-related configurations
	}

	return nil
}

func GetGithubConfig() *GithubConfig {
	return githubConfig
}

type Client struct {
	*github.Client
}

func NewClient() *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "GITHUB_ACCESS_TOKEN"})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return &Client{client}
}
