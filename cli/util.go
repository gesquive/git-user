package cli

import "fmt"
import "strings"
import "os/user"

// ExpandHomeDir expands ~ in a path to the users home directory
func ExpandHomeDir(path string) (expandedPath string) {
	usr, _ := user.Current()
	expandedPath = strings.Replace(path, "~", usr.HomeDir, 1)
	return
}

// ShortenHomeDir replaces the home directory with ~
func ShortenHomeDir(path string) (shortenedPath string) {
	usr, _ := user.Current()
	fmt.Printf("%s %s\n", usr.HomeDir, path)
	shortenedPath = strings.Replace(path, usr.HomeDir, "~", 1)
	return
}
