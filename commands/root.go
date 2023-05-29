package commands

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hbgstcl",
		Short: "Github Stats CLI App",
		Long:  `Github Stats CLI App`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
