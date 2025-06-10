/*
Copyright © 2025 NAME HERE <ekonuma12@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	clientId     string
	clientSecret string
	cmdType      string
	token        string
	host         string = "https://todoist.com"
)

var todoistCmd = &cobra.Command{
	Use:   "todoist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch cmdType {
		case "c":
			configure()
		default:
			fmt.Println("Invalid command type. Use 'c' for configuration.")
		}

	},
}

func configure() {
	fmt.Println("Configuring Todoist integration...")
	var open *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		open = exec.Command("rundll32", "url.dll,FileProtocolHandler", getUrl())
	case "darwin":
		open = exec.Command("open", getUrl())
	case "linux":
		if isWSL() {
			open = exec.Command("wslview", getUrl())
		} else {
			open = exec.Command("xdg-open", getUrl())
		}
	}
	err := open.Start()
	if err != nil {
		fmt.Println("Erro ao abrir o navegador:", err)
	}
	fmt.Println("Please visit the URL in your browser to authorize Todoist integration.")
}

func isWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "Microsoft")
}

func init() {
	rootCmd.AddCommand(todoistCmd)
	todoistCmd.Flags().StringVarP(&cmdType, "configure", "c", "c", "Configuration Flag for Todoist integration")
	todoistCmd.Flags().StringVarP(&clientId, "client-id", "i", "", "Client ID for Todoist integration")
	todoistCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "Client Secret for Todoist integration")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todoistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todoistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetToken(t string) {
	token = t
}

func GetToken() string {
	return token
}

func getUrl() string {
	return fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=data:read_write&state=123&redirect_uri=localhost:8080", host, clientId)
}
