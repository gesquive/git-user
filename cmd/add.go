package cmd

import (
	"os"

	"github.com/gesquive/cli"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add PROFILE_NAME USER_NAME EMAIL",
	Aliases:   []string{"a"},
	Short:     "Add a new profile",
	Long:      `Adds a new user profile that can be used in projects to your config.`,
	ValidArgs: []string{"PROFILE", "USER", "EMAIL"},
	Run:       addRun,
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Usage()
		os.Exit(3)
	}
	name := args[0]
	user := args[1]
	email := args[2]
	cli.Debug("Adding profile %s: '%s' <%s>", name, user, email)
	userProfileConfig.AddProfile(name, user, email)
	cli.Info("Added profile '%s'", name)
}
