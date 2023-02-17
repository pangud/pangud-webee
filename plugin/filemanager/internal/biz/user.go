package biz

type User struct {
}

func (u *User) TableName() string {
	return "ftp_users"
}
