package process

import (
	"encoding/json"
	"fmt"
	"pro05/commom/message"
	"pro05/commom/utils"
)

func NotifyOthersOnlineUser(userId int) {
	var ts utils.Transfer

	var NotifyMes message.NotifyResMsg
	NotifyMes.UserName = AllOnlineInfo[userId].UserName
	NotifyMes.UserId = AllOnlineInfo[userId].UserId
	data, err := json.Marshal(NotifyMes)
	if err != nil {
		fmt.Println("上線通知序列化失敗", err)
	}
	for i, v := range AllOnlineInfo {
		if i == userId {
			continue
		}
		ts.Conn = v.Conn
		ts.WritePkg(data, message.NotifyResMsgType)
	}
}
func NotifyOthersOfflineUser(UserId int) {

	var ts utils.Transfer
	var NotifyMes message.LogOutMsg
	delete(AllOnlineInfo, UserId)
	NotifyMes.UserId = UserId
	data, err := json.Marshal(NotifyMes)
	if err != nil {
		fmt.Println("下線通知序列化失敗", err)
	}
	for _, v := range AllOnlineInfo {
		ts.Conn = v.Conn
		ts.WritePkg(data, message.LogOutMsgType)
	}
}
