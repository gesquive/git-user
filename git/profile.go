package git

import "fmt"

// Profile defines a git profile
type Profile struct {
	Name  string
	User  string
	Email string
}

func (p Profile) String() (out string) {
	out = fmt.Sprintf("%s <%s>", p.User, p.Email)
	return
}

// IsEmpty returns true if this profile has no info
func (p Profile) IsEmpty() bool {
	return len(p.Name) == 0 && len(p.User) == 0 && len(p.Email) == 0
}
