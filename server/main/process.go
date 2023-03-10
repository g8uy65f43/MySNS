package main

import (
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
	"pro05/server/process"
)

type Process struct {
	Conn net.Conn
}

var serverUserProcess *process.UserProcess = &process.UserProcess{}
var SmsProcess *process.SmsProcess = &process.SmsProcess{}

func (thisF *Process) centerProcess(msg *message.MessageData) {
	switch msg.MsgType {
	case message.LoginMsgType:
		serverUserProcess.ServerProcessLogin(*msg)
	case message.RegisterMsgType:
		serverUserProcess.ServerProcessRegister(*msg)
	case message.AddFriendMsgType:
		serverUserProcess.ServerProcessAddFriend(*msg)
	case message.SmsMsgType:
		SmsProcess.SendMessage(*msg)
	case message.OffLineMsgType:
		SmsProcess.GetOffLineMsg(*msg)
	case message.GSmsMsgType:
		SmsProcess.GSendMessage(*msg)

	}

}

func (thisF *Process) process() (err error) {
	tf := utils.Transfer{
		Conn: thisF.Conn,
	}
	var LogOutIndex int
	serverUserProcess.Conn = thisF.Conn
	SmsProcess.Conn = thisF.Conn
	msg, err := tf.ReadPkg()
	if err != nil {
		for _, v := range process.AllOnlineInfo {
			if v.Conn == thisF.Conn {
				LogOutIndex = v.UserId
				break
			}
		}
		if LogOutIndex != 0 {
			for _, v := range process.AllOnlineInfo {
				if v.UserId == LogOutIndex {
					continue
				}
				process.NotifyOthersOfflineUser(LogOutIndex)
			}
		}
		fmt.Println("伺服器端獲取資料失敗", err)
	}
	thisF.centerProcess(&msg)
	return
}
