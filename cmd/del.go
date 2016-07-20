package cmd

import (
	"os"

	cli "github.com/gesquive/cli-log"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:       "del PROFILE_NAME",
	Aliases:   []string{"d", "delete"},
	Short:     "Delete a profile",
	Long:      `Delete a user profile from your config.`,
	ValidArgs: []string{"PROFILE"},
	Run:       delRun,
}

func init() {
	RootCmd.AddCommand(delCmd)
}

func delRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(3)
	}
	name := args[0]
	cli.Debug("Deleting profile %s", name)
	userProfileConfig.DeleteProfile(name)
	cli.Info("Deleted profile '%s'", name)
}
