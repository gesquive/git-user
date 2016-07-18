package git

import "strings"
import "github.com/codeskyblue/go-sh"

var git = "git"

// SetGitPath will set a custom git path if necessary
func SetGitPath(gitPath string) {
	git = gitPath
}

// Exists checks to make sure it can find the git executable
func Exists() bool {
	out, err := sh.Command("which", git).Output()
	if err == nil && len(out) > 0 {
		return true
	}
	return false
}

// Version returns the git version in use
func Version() string {
	version := ""
	out, err := sh.Command(git, "--version").Output()
	if err != nil || len(out) == 0 {
		return ""
	}
	parts := strings.Split(string(out), " ")
	if len(parts) != 3 {
		return ""
	}
	if parts[0] != "git" || parts[1] != "version" {
		return ""
	}
	version = strings.Trim(parts[2], "\n\r ")
	return version
}

// GetUser returns the global git user info
func GetUser() (user string, email string) {
	out, err := sh.Command(git, "config", "user.name").Output()
	if err == nil {
		user = strings.TrimSpace(string(out))
	}

	out, err = sh.Command(git, "config", "user.email").Output()
	if err == nil {
		email = strings.TrimSpace(string(out))
	}
	return
}

// SetGlobalUser sets the user for the global git config
func SetGlobalUser(user string, email string) {
	sh.Command(git, "config", "--global", "user.name", user).Run()
	sh.Command(git, "config", "--global", "user.email", email).Run()
}

// RemoveGlobalUser removes user from the global git config
func RemoveGlobalUser() {
	sh.Command(git, "config", "--global", "--remove-section", "user").Run()
}
