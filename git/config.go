package git

import "fmt"
import "os"
import "strings"
import "github.com/codeskyblue/go-sh"
import "github.com/go-ini/ini"
import "github.com/gesquive/git-user/cli"

// UserProfileConfig is the config file we want to edit
type UserProfileConfig struct {
	path    string
	iniFile *ini.File
}

// NewUserProfileConfig initializes and returns a new UserProfileConfig
func NewUserProfileConfig(hintPath string) (*UserProfileConfig, error) {
	defaultProfilePath := cli.ExpandHomeDir("~/.git_profiles")
	config := new(UserProfileConfig)

	if _, err := os.Stat(hintPath); err == nil {
		config.path = cli.ExpandHomeDir(hintPath)
	} else {
		out, err := sh.Command("git", "config", "user.profiles").Output()
		if err == nil && len(out) > 0 {
			config.path = cli.ExpandHomeDir(strings.TrimSpace(string(out)))
		} else {
			// Then there is no path in the gitconfig, set it
			sh.Command("git", "config", "--global", "user.profiles", defaultProfilePath)
			config.path = defaultProfilePath
		}
	}

	if len(config.path) == 0 {
		err := fmt.Errorf("No config path could be found.")
		return nil, err
	}

	var iniFile *ini.File
	if _, err := os.Stat(config.path); os.IsNotExist(err) {
		// Config file doesn't exist, create it
		iniFile = ini.Empty()
	} else {
		iniFile, err = ini.Load(config.path)
		if err != nil {
			return nil, err
		}
	}

	config.iniFile = iniFile

	return config, nil
}

func (u UserProfileConfig) save() (err error) {
	err = u.iniFile.SaveTo(u.path)
	return
}

// Path returns the user profile path
func (u UserProfileConfig) Path() string {
	return u.path
}

// CheckProfile checks if name exists as a profile
func (u UserProfileConfig) CheckProfile(name string) bool {
	profiles := u.iniFile.SectionStrings()
	for _, profileName := range profiles {
		if strings.EqualFold(name, profileName) {
			return true
		}
	}
	return false
}

// AddProfile adds a profile
func (u UserProfileConfig) AddProfile(name string, user string, email string) {
	profileSection := u.iniFile.Section(name)
	profileSection.NewKey("name", user)
	profileSection.NewKey("email", email)
	u.save()
}

// EditProfile edits a profile
func (u UserProfileConfig) EditProfile(name string, user string, email string) {
	profileSection := u.iniFile.Section(name)
	profileSection.NewKey("name", user)
	profileSection.NewKey("email", email)
	u.save()
}

// DeleteProfile removes a profile from the config
func (u UserProfileConfig) DeleteProfile(name string) {
	u.iniFile.DeleteSection(name)
	u.save()
}

// GetProfile gets the profile info
func (u UserProfileConfig) GetProfile(name string) (profile Profile) {
	profileSection, err := u.iniFile.GetSection(name)
	if err == nil {
		userKey, err := profileSection.GetKey("name")
		if err == nil {
			profile.User = userKey.String()
		}
		emailKey, err := profileSection.GetKey("email")
		if err == nil {
			profile.Email = emailKey.String()
		}
		if !profile.IsEmpty() {
			profile.Name = name
		}
	}
	return
}

// GetAllProfiles gets all of the saved profiles
func (u UserProfileConfig) GetAllProfiles() (profiles []Profile) {
	profileSections := u.iniFile.Sections()

	for _, profileSection := range profileSections {
		profile := Profile{}
		profile.Name = profileSection.Name()
		userKey, err := profileSection.GetKey("name")
		if err == nil {
			profile.User = userKey.String()
		}
		emailKey, err := profileSection.GetKey("email")
		if err == nil {
			profile.Email = emailKey.String()
		}
		if len(profile.Email) > 0 && len(profile.User) > 0 {
			profiles = append(profiles, profile)
		}
	}
	return profiles
}
