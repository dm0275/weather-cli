package version

import (
	"fmt"
	"github.com/dm0275/weather-cli/pkg/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print WeatherCLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Version())
		},
	}

	return versionCmd
}
