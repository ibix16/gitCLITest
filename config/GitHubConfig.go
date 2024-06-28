package config

import (
	"github.com/spf13/viper"
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
