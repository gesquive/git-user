// Copyright © 2016 Gus Esquivel <gesquive@gmail.com>

package cmd

import (
	"os"

	"github.com/gesquive/git-user/cli"
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
	cli.Debugf("Deleting profile %s", name)
	userProfileConfig.DeleteProfile(name)
	cli.Infof("Deleted profile '%s'", name)
}
