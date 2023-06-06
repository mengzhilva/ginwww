package modules

type User struct {
	Id   int
	Name string
}

func (User) TableName() string {
	return "go_user"
}
