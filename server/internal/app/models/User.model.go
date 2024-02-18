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
	Tokens   []Token `json:"-"`
	Role     Role `json:"role"`
	CreatedGroups []Group `gorm:"foreignKey:CreatorID" json:"-"` 
	Groups []*Group `gorm:"many2many:user_groups;" json:"-"` 
	ReceivedMessages []PrivateMessage `gorm:"foreignKey:ReceiverId" json:"-"`
	SentMessages []PrivateMessage `gorm:"foreignKey:SenderId"  json:"-"`
}

