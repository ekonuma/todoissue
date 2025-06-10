/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubToken string

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("github called")
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.Flags().StringVarP(&githubToken, "client-id", "a", "", "Enable Github integration")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
