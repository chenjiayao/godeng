package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var VERSION = "v1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of godeng",
	Long:  `All software has versions. This is godeng's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("godeng ", VERSION)
	},
}
