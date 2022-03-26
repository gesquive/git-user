package main

import "github.com/gesquive/git-user/cmd"

var (
	buildVersion = "v2.0.6-dev"
	buildCommit  = ""
	buildDate    = ""
)

func main() {
	cmd.BuildVersion = buildVersion
	cmd.BuildCommit = buildCommit
	cmd.BuildDate = buildDate
	cmd.Execute()
}
