// Copyright © 2016 Gus Esquivel <gesquive@gmail.com>

package cmd

import (
	"github.com/gesquive/git-user/cli"
	"github.com/spf13/cobra"
	"os"
)

var editCmd = &cobra.Command{
	Use:       "edit PROFILE_NAME USER_NAME EMAIL",
	Aliases:   []string{"e"},
	Short:     "Edit a profile",
	Long:      `Edit a user profile name or email address in your config.`,
	ValidArgs: []string{"PROFILE", "USER", "EMAIL"},
	Run:       editRun,
}

func init() {
	RootCmd.AddCommand(editCmd)
}

func editRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Usage()
		os.Exit(3)
	}
	name := args[0]
	user := args[1]
	email := args[2]
	cli.Debugf("Editing profile %s: '%s' <%s>", name, user, email)
	userProfileConfig.AddProfile(name, user, email)
	cli.Infof("Edited profile '%s'", name)
}
