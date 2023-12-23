package cmd

import (
	"github.com/dm0275/weather-cli/cmd/city"
	"github.com/dm0275/weather-cli/cmd/version"
	"github.com/spf13/cobra"
)

type WeatherCLI struct {
	rootCmd *cobra.Command
}

func NewCLI() *WeatherCLI {
	weatherCli := &WeatherCLI{}
	weatherCli.initialize()

	// Setup sub-commands
	weatherCli.rootCmd.AddCommand(version.NewVersionCmd())
	weatherCli.rootCmd.AddCommand(city.NewCityCmd())

	// Disable auto-completion
	weatherCli.rootCmd.CompletionOptions.DisableDefaultCmd = true

	return weatherCli
}

func (w *WeatherCLI) initialize() {
	// Initialize the root cmd
	w.rootCmd = &cobra.Command{
		Use:   "weather-cli",
		Short: "weather-cli - a CLI to get weather data",
		Long:  `weather-cli is CLI used to get weather data`,
	}
}

func (w *WeatherCLI) Execute() {
	cobra.CheckErr(w.rootCmd.Execute())
}
