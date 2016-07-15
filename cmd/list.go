// Copyright © 2016 Gus Esquivel <gesquive@gmail.com>

package cmd

import (
	"github.com/gesquive/git-user/cli"
	"github.com/gesquive/git-user/git"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all saved profiles",
	Long:    `List all of the saved user profiles found in your config.`,
	Run:     listRun,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func listRun(cmd *cobra.Command, args []string) {
	if gitRepo.HasUserSet() {
		cli.Infof("Project Profile:")
		cli.Infof("  Path: %s", cli.Green(projectPath))
	} else {
		cli.Infof("Global Profile:")
	}
	user, email := git.GetUser()
	if len(user) == 0 {
		user = cli.Red("N/A")
	}
	if len(email) == 0 {
		email = cli.Red("N/A")
	}
	cli.Infof("  User: %s <%s>", cli.Green(user), cli.Blue(email))
	cli.Infof("")

	profiles := userProfileConfig.GetAllProfiles()
	if len(profiles) == 0 {
		cli.Infof("There are no profiles in your config.")
		cli.Infof("  Add a profile with \"%s add <profile> <name> <email>\"",
			"git user")
		cli.Infof("Type \"{} --help\" for more info.")
	} else {
		cli.Infof("Saved Profiles:")
		for _, profile := range profiles {
			cli.Infof("  %s: %s", cli.Yellow(profile.Name), profile.String())
		}
	}
}
