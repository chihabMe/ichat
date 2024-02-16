package models

type Role int

const (
	NormalUserRole Role = iota
	AdminUserRole
)

var roleStrings = [...]string{
	"normal",
	"admin",
}

func (r Role) String() string {
	if r < NormalUserRole || r > AdminUserRole {
		return "unknown"
	}
	return roleStrings[r]
}

type User struct {
	Base
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Verified bool `json:"verified"`
	Active   bool `json:"active"`
	Profile  Profile `gorm:"foreignKey:UserId" json:"profile"`
	Tokens   []Token
	Role     Role `json:"role"`
}

