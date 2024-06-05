package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (*User) TableName() string {
	return "user"
}
