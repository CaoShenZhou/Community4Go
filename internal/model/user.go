package Model

type User struct {
	*Model
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (u User) TableName() string {
	return "user"
}
