package cmd

import (
	"github.com/spf13/cobra"
)

type WeatherCLI struct {
	rootCmd *cobra.Command
}

func NewCLI() *WeatherCLI {
	weatherCli := &WeatherCLI{}
	weatherCli.initialize()

	return weatherCli
}

func (c WeatherCLI) initialize() {
	// Initialize the root cmd
	c.rootCmd = &cobra.Command{
		Use:   "weather-cli",
		Short: "weather-cli - a CLI to get weather data",
		Long:  `weather-cli is CLI used to get weather data`,
	}
}

func (c WeatherCLI) Execute() {
	cobra.CheckErr(c.rootCmd.Execute())
}
