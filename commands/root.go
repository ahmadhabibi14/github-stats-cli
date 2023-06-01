package commands

import (
	"github.com/ahmadhabibi14/github-stats-cli/app"
	"github.com/spf13/cobra"
)

var username string
var rootCmd = &cobra.Command{
	Use:   "hbgstcl",
	Short: "Github Stats CLI App",
	Long:  `Github Stats CLI App`,
	Run: func(cmd *cobra.Command, args []string) {
		app.MainApp(username)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Github username")
	rootCmd.MarkFlagRequired("username")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
