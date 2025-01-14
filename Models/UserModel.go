package Models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsDelete bool   `json:"isDelete" gorm:"column:IsDelete"`
}

func (b *User) TableName() string {
	return "dbo.user"
}
