package biz

import "github.com/pangud/pangud/pkg/errors"

const (

	// 以1开头为系统端错误 1  01 标识 core模块 code begin with 001
	ErrCodeSaveUserError errors.ErrorCode = 101001
	ErrCodeUserNotFound  errors.ErrorCode = 101002
	/*以2开头为用户错误 01 标识 core模块
	200000-200999 用户错误*/
	ErrCodeUsernameOrPasswordIsWrong errors.ErrorCode = 201001
	ErrCodeUserInfoIsInvalid         errors.ErrorCode = 201002

	// 以3开头为外部系统错误 3  00 标识 core模块
)

var (
	ErrSaveUserError             = errors.New(ErrCodeSaveUserError, "save user erro")
	ErrUserNotFound              = errors.New(ErrCodeUserNotFound, "user not found")
	ErrUsernameOrPasswordIsWrong = errors.New(ErrCodeUsernameOrPasswordIsWrong, "username or password is wrong")
	ErrUserInfoIsInvalid         = errors.New(ErrCodeUserInfoIsInvalid, "user info is invalid")
)
