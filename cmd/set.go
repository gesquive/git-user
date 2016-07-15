// Copyright © 2016 Gus Esquivel <gesquive@gmail.com>

package cmd

import (
	"github.com/gesquive/git-user/cli"
	"github.com/gesquive/git-user/git"
	"github.com/spf13/cobra"
	"os"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:       "set PROFILE_NAME",
	Aliases:   []string{"s"},
	Short:     "Set the profile for the current project",
	Long:      `Set the default user profile for the current project.`,
	ValidArgs: []string{"PROFILE"},
	Run:       setRun,
}

func init() {
	RootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().BoolVarP(&global, "global", "g", false,
		"Apply the profile to the global config")

}

func setRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(3)
	}
	name := args[0]
	cli.Debugf("Setting profile %s", name)
	profile := userProfileConfig.GetProfile(name)
	if profile.IsEmpty() {
		cli.Infof("There is no profile named '%s' in the config", name)
		cli.Infof("You can add the profile with:")
		//TODO: replace the app name with the detected app name
		cli.Infof("  '%s add %s <name> <email>'", "git-user", "name")
	} else {
		if global {
			git.SetGlobalUser(profile.User, profile.Email)
			cli.Infof("The global user has been set too '%s <%s>'",
				profile.User, profile.Email)
		} else {
			gitRepo.SetUser(profile.User, profile.Email)
			cli.Infof("The user for the '%s' repository has been set too '%s <%s>'",
				gitRepo.Name(), profile.User, profile.Email)
		}
	}
}
