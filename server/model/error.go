package model

//error
import "errors"

var (
	ERROR_USER_NOTEXISTS   = errors.New("用戶不存在...")
	ERROR_USER_EXISTS      = errors.New("用戶已存在...")
	ERROR_FRIEND_EXISTS    = errors.New("好友已存在...")
	ERROR_USER_PWD         = errors.New("密碼不正確...")
	ERROR_USER_STYLE_ERROR = errors.New("格式錯誤...")
)
