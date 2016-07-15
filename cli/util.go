package cli

import "strings"
import "os/user"

// ExpandHomeDir expands ~ in a path to the users home directory
func ExpandHomeDir(path string) (expandedPath string) {
	usr, _ := user.Current()
	expandedPath = strings.Replace(path, "~", usr.HomeDir, 1)
	return
}
