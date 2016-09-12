package cmd

import (
	cli "github.com/gesquive/cli-log"
	"github.com/gesquive/git-user/git"
	"github.com/spf13/cobra"
	"path/filepath"
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
		cli.Info("Project Profile:")
		fullPath, err := filepath.Abs(filepath.Clean(projectPath))
		if err != nil {
			fullPath = projectPath
		}
		cli.Info("  Path: %s", cli.Green(fullPath))
	} else {
		cli.Info("Global Profile:")
	}
	user, email := git.GetUser()
	if len(user) == 0 {
		user = cli.Red("N/A")
	}
	if len(email) == 0 {
		email = cli.Red("N/A")
	}
	cli.Info("  User: %s <%s>", cli.Green(user), cli.Blue(email))
	cli.Info("")

	profiles := userProfileConfig.GetAllProfiles()
	if len(profiles) == 0 {
		cli.Info("There are no profiles in your config.")
		cli.Info("  Add a profile with \"%s add <profile> <name> <email>\"",
			"git user")
		cli.Info("Type \"{} --help\" for more info.")
	} else {
		cli.Info("Saved Profiles:")
		for _, profile := range profiles {
			cli.Info("  %s: %s", cli.Yellow(profile.Name), profile.String())
		}
	}
}
