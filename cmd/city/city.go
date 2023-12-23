package city

import (
	"github.com/spf13/cobra"
)

type Config struct {
	City string
}

func NewCityCmd() *cobra.Command {
	config := &Config{}
	cityCmd := &cobra.Command{
		Use:   "city",
		Short: "Get weather for City",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Implement

		},
	}

	// Configure flags
	configureFlags(cityCmd, config)

	return cityCmd
}

func configureFlags(cmd *cobra.Command, config *Config) {
	cmd.Flags().StringVarP(&config.City, "city", "", "", "City to fetch weather data")
}
