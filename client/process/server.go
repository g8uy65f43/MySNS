package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
)

func server(conn net.Conn) {

	fmt.Println("持續通訊監聽中")
	ts := utils.Transfer{
		Conn: conn,
	}
	for {
		dataFromServer, err := ts.ReadPkg()
		if err != nil {
			fmt.Println("持續通訊監聽出錯", err)
			return
		}
		switch dataFromServer.MsgType {
		case message.NotifyResMsgType:
			var NotifyResData message.NotifyResMsg
			err := json.Unmarshal([]byte(dataFromServer.Data), &NotifyResData)
			if err != nil {
				fmt.Println("上線訊息反序列化失敗")
			}
			fmt.Println(NotifyResData.UserName, "上線啦!")
			AllOnlineInfo[NotifyResData.UserId] = message.OnlineInfoDemo{UserId: NotifyResData.UserId, UserName: NotifyResData.UserName, UserStatus: 1}
		case message.LogOutMsgType:
			var LogOutResData message.LogOutMsg
			err := json.Unmarshal([]byte(dataFromServer.Data), &LogOutResData)
			if err != nil {
				fmt.Println("下線訊息反序列化失敗")
			}
			fmt.Println(AllOnlineInfo[LogOutResData.UserId].UserName, "下線啦!")
			delete(AllOnlineInfo, LogOutResData.UserId)
		case message.SmsMsgType:
			var SmsMsg message.SmsMsg
			json.Unmarshal([]byte(dataFromServer.Data), &SmsMsg)
			fmt.Printf("收到來自[%v]的私信\t:\t%v\n", SmsMsg.UserName, SmsMsg.Msg)
		case message.GSmsMsgType:
			var GSmsMsg message.GSmsMsg
			json.Unmarshal([]byte(dataFromServer.Data), &GSmsMsg)
			fmt.Printf("收到來自[%v]的群發\t:\t%v\n", GSmsMsg.UserName, GSmsMsg.Msg)
		default:
			fmt.Println("沒找到", dataFromServer.MsgType)
		}
	}
}
