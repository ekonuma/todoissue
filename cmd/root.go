/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	oauthConfig  *oauth2.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-issues",
	Short: "Busca issues no GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Configuração do OAuth
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		tc := oauth2.NewClient(ctx, ts)

		// Criando um cliente para a API do GitHub
		client := github.NewClient(tc)

		// Listando os repositórios da organização "minha-organização"
		opts := &github.RepositoryListOptions{}
		for {
			repos, resp, err := client.Repositories.List(ctx, "ekonuma", opts)
			if err != nil {
				fmt.Printf("Error listing repos: %v", err)
				os.Exit(1)
			}

			for _, repo := range repos {
				fmt.Println(*repo.Name)
				issueOpts := &github.IssueListByRepoOptions{}
				issues, _, err := client.Issues.ListByRepo(ctx, *repo.Owner.Login, *repo.Name, issueOpts)
				if err != nil {
					fmt.Printf("Error listing issues for %s: %v", *repo.Name, err)
					continue
				}

				for _, issue := range issues {
					fmt.Printf("Issue: %s\n", *issue.Title)
					// Processar cada issue
				}
			}

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todoissue.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getIssues(ctx context.Context, client *github.Client) {
	// Itera sobre todos os repositórios do usuário
	opts := &github.RepositoryListOptions{
		Type: "all",
	}
	for {
		repos, resp, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, repo := range repos {
			issues, _, err := client.Issues.ListByRepo(ctx, *repo.Owner.Login, *repo.Name, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, issue := range issues {
				fmt.Printf("%s: %s\n", *issue.Title, *issue.HTMLURL)
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
}
