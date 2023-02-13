package biz

type User struct {
}

func (u *User) TableName() string {
	return "p_file_ftp_users"
}
