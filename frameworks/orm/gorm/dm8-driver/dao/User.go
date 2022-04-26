package dao

type User struct {
	Id        int    `gorm:"primaryKey column:Id"`
	UserName  string `gorm:"UserName"`
	Password  string `gorm:"Password"`
	FileTable string `FileTable:"FileTable"`
}

func NewUser(username string, password string) *User {
	return &User{
		UserName: username,
		Password: password,
	}
}
