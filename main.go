package main

import "os"
import "fmt"
import "path/filepath"
import "github.com/gesquive/git-user/cmd"

var version = "v2.0.1"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("%s %s%s",
		filepath.Base(os.Args[0]),
		version,
		dirty)
	cmd.Execute(displayVersion)
}
