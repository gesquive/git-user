package git

import "path/filepath"
import "strings"
import "github.com/codeskyblue/go-sh"

// Repo allows you to execute actions on a specific git project
type Repo struct {
	path string
}

// NewGitRepo constructs a git repo for you
func NewGitRepo(path string) *Repo {
	repo := new(Repo)
	repo.path = path
	return repo
}

// Path returns the repos full path
func (r Repo) Path() string {
	return r.path
}

// Name returns the name of the repo, this is just the base folder name
func (r Repo) Name() string {
	absPath := r.path
	absPath, err := filepath.Abs(r.path)
	if err != nil {
		absPath = r.path
	}
	return filepath.Base(absPath)
}

// HasUserSet returns if repo has a user set
func (r Repo) HasUserSet() bool {
	value, err := sh.Command(git, "-C", r.path, "config",
		"--local", "user.name").Output()
	if err == nil && len(value) > 0 {
		return true
	}
	return false
}

// GetUser returns the repo user
func (r Repo) GetUser() (user string, email string) {
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

// SetUser sets the user for a git repo
func (r Repo) SetUser(user string, email string) {
	sh.Command(git, "-C", r.path, "config", "user.name", user).Run()
	sh.Command(git, "-C", r.path, "config", "user.email", email).Run()
}

// RemoveUser removes the set user from the git repo
func (r Repo) RemoveUser() {
	sh.Command(git, "-C", r.path, "config", "--remove-section", "user").Run()
}
