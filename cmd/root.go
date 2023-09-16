package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "weather-cli",
		Short: "weather-cli - a CLI to get weather data",
		Long:  `weather-cli is CLI used to get weather data`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Add cmd action and subcommands
			fmt.Println("TODO: Add subcommands")
		},
	}
}

func Execute() error {
	return rootCmd().Execute()
}
