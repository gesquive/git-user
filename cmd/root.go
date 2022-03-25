package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gesquive/cli"
	"github.com/gesquive/git-user/git"
	"github.com/gesquive/git-user/user"
	"github.com/spf13/cobra"
)

// current build info
var (
	BuildVersion = "v0.1.0-dev"
	BuildCommit  = ""
	BuildDate    = ""
)

var userProfileConfig *git.UserProfileConfig
var gitRepo *git.Repo
var gitPath string
var cfgFilePath string
var projectPath string
var appName string

var logDebug bool
var showVersion bool

var useGlobal bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:              "git-user",
	Short:            "Allows you to save multiple user profiles and set them as git project defaults",
	Long:             `git-user lets you quickly switch between multiple git user profiles`,
	PersistentPreRun: persistentPreRun,
	Run:              listRun,
	Hidden:           true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	RootCmd.SetHelpTemplate(helpTemplate())
	RootCmd.SetUsageTemplate(usageTemplate())
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//TODO: Perform some sort of ENV var expansion to help with documentation and help
	defaultConfigPath := "~/.git-profiles"
	defaultProjectPath := "."

	RootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c",
		user.ShortenHomeDir(defaultConfigPath), "config file")
	RootCmd.PersistentFlags().StringVarP(&projectPath, "path", "p",
		user.ShortenHomeDir(defaultProjectPath), "The project to get/set the user")
	RootCmd.PersistentFlags().StringVarP(&gitPath, "git-path", "g", "git",
		"The git executable to use")

	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false,
		"Show the version and exit")
	RootCmd.PersistentFlags().BoolVarP(&logDebug, "debug", "D", false,
		"Write debug messages to console")
	RootCmd.PersistentFlags().MarkHidden("debug")

	// TODO: activate viper for ENV vars
	// viper.SetEnvPrefix("git-user")
	// viper.AutomaticEnv()
}

func initConfig() {
	if logDebug {
		cli.SetPrintLevel(cli.LevelDebug)
	}
	cli.Debug("Running with debug turned on")

	appName = os.Args[0]
	if appName == "git-user" {
		appName = "git user"
	}

	git.SetGitPath(gitPath)
	if !git.Exists() {
		cli.Info("Could not find a valid git executable")
		cli.Errorf("'%s' was not found\n", gitPath)
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

	cfgFilePath = os.ExpandEnv(user.ExpandHomeDir(cfgFilePath))
	projectPath = os.ExpandEnv(user.ExpandHomeDir(projectPath))
	cli.Debug("configPath='%s'", cfgFilePath)
	cli.Debug("projectPath='%s'", projectPath)
	var err error
	userProfileConfig, err = git.NewUserProfileConfig(cfgFilePath)
	if err != nil {
		cli.Errorf("%v\n", err)
		os.Exit(2)
	}
	cli.Debug("profileConfigPath=%s", userProfileConfig.Path())
	gitRepo = git.NewGitRepo(projectPath)
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	if showVersion {
		fmt.Printf("github.com/gesquive/git-user\n")
		fmt.Printf(" Version:    %s\n", BuildVersion)
		if len(BuildCommit) > 6 {
			fmt.Printf(" Git Commit: %s\n", BuildCommit[:7])
		}
		if BuildDate != "" {
			fmt.Printf(" Build Date: %s\n", BuildDate)
		}
		fmt.Printf(" Go Version: %s\n", runtime.Version())
		fmt.Printf(" OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}
	if logDebug {
		cli.SetPrintLevel(cli.LevelDebug)
	}
}

func helpTemplate() string {
	return fmt.Sprintf("%s\nVersion:\n  github.com/gesquive/git-user %s\n",
		RootCmd.HelpTemplate(), BuildVersion)
}

func usageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
  
Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}
  
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
  
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
  
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
  
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
