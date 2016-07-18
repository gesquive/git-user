package cmd

import (
	"fmt"
	"os"

	"github.com/gesquive/git-user/cli"
	"github.com/gesquive/git-user/git"
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
		cli.ShortenHomeDir(defaultConfigPath), "config file")
	RootCmd.PersistentFlags().StringVarP(&projectPath, "path", "p",
		cli.ShortenHomeDir(defaultProjectPath), "The project to get/set the user")
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
		cli.PrintDebug = true
	}
	if showVersion {
		cli.Infof(displayVersion)
		os.Exit(0)
	}
	cli.Debugf("Running with debug turned on")

	appName = os.Args[0]
	if appName == "git-user" {
		appName = "git user"
	}

	git.SetGitPath(gitPath)
	if !git.Exists() {
		cli.Infof("Could not find a valid git executable")
		cli.Errorf("'%s' was not found", gitPath)
		os.Exit(5)
	}
	gitVersion := git.Version()
	if len(gitVersion) == 0 {
		cli.Infof("Git version is not valid")
		cli.Errorf("The git executable found might not be valid")
		os.Exit(6)
	}
	cli.Debugf("gitPath=%s", gitPath)
	cli.Debugf("gitVersion=%s", gitVersion)

	cfgFilePath = cli.ExpandHomeDir(cfgFilePath)
	projectPath = cli.ExpandHomeDir(projectPath)
	cli.Debugf("configPath='%s'", cfgFilePath)
	cli.Debugf("projectPath='%s'", projectPath)
	var err error
	userProfileConfig, err = git.NewUserProfileConfig(cfgFilePath)
	if err != nil {
		cli.Errorf("%v", err)
		os.Exit(2)
	}
	cli.Debugf("profileConfigPath=%s", userProfileConfig.Path())
	gitRepo = git.NewGitRepo(projectPath)
}
