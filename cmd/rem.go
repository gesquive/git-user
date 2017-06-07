package cmd

import (
	"github.com/gesquive/cli"
	"github.com/gesquive/git-user/git"
	"github.com/spf13/cobra"
	"os"
)

var remCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"r", "rem", "remove"},
	Short:   "Remove a profile from the current project",
	Long:    `Remove a user profile from the default profile for a project.`,
	Run:     remRun,
}

func init() {
	RootCmd.AddCommand(remCmd)
	remCmd.PersistentFlags().BoolVarP(&global, "global", "G", false,
		"Remove the profile from the global config")
}

func remRun(cmd *cobra.Command, args []string) {

	if len(args) == 1 {
		cli.Info(cli.Yellow("If you are trying to delete a profile, use the 'del' command.\n"))
	}
	if len(args) > 0 {
		cmd.Usage()
		os.Exit(3)
	}

	if global {
		cli.Debug("Removing global user")
		git.RemoveGlobalUser()
		cli.Info("Removed user info from the global config")
	} else {
		cli.Debug("Removing project user")
		gitRepo.RemoveUser()
		cli.Info("Removed user info from '%s'", gitRepo.Name())
	}
}
