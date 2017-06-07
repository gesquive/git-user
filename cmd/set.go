package cmd

import (
	"github.com/gesquive/cli"
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
	setCmd.PersistentFlags().BoolVarP(&global, "global", "G", false,
		"Apply the profile to the global config")

}

func setRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(3)
	}
	name := args[0]
	cli.Debug("Setting profile %s", name)
	profile := userProfileConfig.GetProfile(name)
	if profile.IsEmpty() {
		cli.Info("There is no profile named '%s' in the config", name)
		cli.Info("You can add the profile with:")
		cli.Info("  '%s add %s NAME EMAIL'", appName, name)
	} else {
		if global {
			git.SetGlobalUser(profile.User, profile.Email)
			cli.Info("The global user has been set to '%s <%s>'",
				profile.User, profile.Email)
		} else {
			gitRepo.SetUser(profile.User, profile.Email)
			cli.Info("The user for the '%s' repository has been set to '%s <%s>'",
				gitRepo.Name(), profile.User, profile.Email)
		}
	}
}
