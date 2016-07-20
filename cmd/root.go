package cmd

import (
	"fmt"
	"os"

	cli "github.com/gesquive/cli-log"
	"github.com/gesquive/git-user/git"
	"github.com/gesquive/git-user/user"
	"github.com/spf13/cobra"
)

var userProfileConfig *git.UserProfileConfig
var gitRepo *git.Repo
var gitPath string
var cfgFilePath string
var projectPath string
var displayVersion string
var appName string

var logDebug bool
var showVersion bool

var global bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "git-user",
	Short: "Allows you to save multiple user profiles and set them as git project defaults",
	Long:  `git-user lets you quickly switch between multiple git user profiles`,
	Run:   listRun,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	displayVersion = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	defaultConfigPath := "~/.git-profiles"
	defaultProjectPath, _ := os.Getwd()

	RootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c",
		user.ShortenHomeDir(defaultConfigPath), "config file")
	RootCmd.PersistentFlags().StringVarP(&projectPath, "path", "p",
		user.ShortenHomeDir(defaultProjectPath), "The project to get/set the user")
	RootCmd.PersistentFlags().StringVarP(&gitPath, "git-path", "g", "git",
		"The git executable to use")
	RootCmd.PersistentFlags().BoolVarP(&logDebug, "debug", "D", false,
		"Write debug messages to console")
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "V", false,
		"Show the version and exit")

	RootCmd.PersistentFlags().MarkHidden("debug")
}

func initConfig() {
	if logDebug {
		cli.SetLogLevel(cli.LevelDebug)
	}
	if showVersion {
		cli.Info(displayVersion)
		os.Exit(0)
	}
	cli.Debug("Running with debug turned on")

	appName = os.Args[0]
	if appName == "git-user" {
		appName = "git user"
	}

	git.SetGitPath(gitPath)
	if !git.Exists() {
		cli.Info("Could not find a valid git executable")
		cli.Error("'%s' was not found", gitPath)
		os.Exit(5)
	}
	gitVersion := git.Version()
	if len(gitVersion) == 0 {
		cli.Info("Git version is not valid")
		cli.Error("The git executable found might not be valid")
		os.Exit(6)
	}
	cli.Debug("gitPath=%s", gitPath)
	cli.Debug("gitVersion=%s", gitVersion)

	cfgFilePath = user.ExpandHomeDir(cfgFilePath)
	projectPath = user.ExpandHomeDir(projectPath)
	cli.Debug("configPath='%s'", cfgFilePath)
	cli.Debug("projectPath='%s'", projectPath)
	var err error
	userProfileConfig, err = git.NewUserProfileConfig(cfgFilePath)
	if err != nil {
		cli.Error("%v", err)
		os.Exit(2)
	}
	cli.Debug("profileConfigPath=%s", userProfileConfig.Path())
	gitRepo = git.NewGitRepo(projectPath)
}
