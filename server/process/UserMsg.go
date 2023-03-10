package process

import (
	"net"
)

type ServerOnlineInfo struct {
	UserId     int
	Conn       net.Conn
	UserStatus int
	UserName   string
}

var AllOnlineInfo map[int]ServerOnlineInfo = make(map[int]ServerOnlineInfo)
