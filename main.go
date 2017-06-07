package main

import "fmt"
import "github.com/gesquive/git-user/cmd"

var version = "v2.0.5"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("git-user %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
