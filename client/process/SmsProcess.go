package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
)

type SmsProcess struct {
}

func (thisF *SmsProcess) SendMessage(otherId int, msg string) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("發送訊息錯誤", err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var SmsMsg message.SmsMsg = message.SmsMsg{UserId: MyId, OtherId: otherId, Msg: msg}
	data, err := json.Marshal(SmsMsg)
	if err != nil {
		fmt.Println("發送訊息序列化失敗", err)
		return
	}
	tf.WritePkg(data, message.SmsMsgType)

	var SmsResMsg message.SmsResMsg
	MsgData, err := tf.ReadPkg()
	json.Unmarshal([]byte(MsgData.Data), &SmsResMsg)
	if SmsResMsg.ResCode == 200 {
		fmt.Println("傳送成功!")
	} else if SmsResMsg.ResCode == 220 {
		fmt.Println("傳送成功，但對方不在線上，轉換為留言。")
	} else if SmsResMsg.ResCode == 300 {
		fmt.Println(SmsResMsg.Error)
	} else if SmsResMsg.ResCode == 400 {
		fmt.Println(SmsResMsg.Error)
	} else {
		fmt.Println("傳送訊息未知的錯誤", err)
	}
}
func (thisF *SmsProcess) GetOffLineMessage() {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("發送訊息錯誤", err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var OffLineMsg message.OffLineMsg = message.OffLineMsg{UserId: MyId}
	data, err := json.Marshal(OffLineMsg)
	if err != nil {
		fmt.Println("序列化獲取離線留言請求失敗")
		return
	}
	tf.WritePkg(data, message.OffLineMsgType)

	var OffLineResMsg message.OffLineResMsg
	ResData, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("獲取離線留言失敗")
		return
	}
	json.Unmarshal([]byte(ResData.Data), &OffLineResMsg)
	if OffLineResMsg.ResCode == 200 {
		fmt.Println("成功獲取留言")
		for _, v := range OffLineResMsg.Msgs {
			var getOfflineMsg message.SmsMsg
			json.Unmarshal([]byte(v), &getOfflineMsg)
			fmt.Printf("收到來自[%v]的訊息\t:\t%v\n", getOfflineMsg.UserName, getOfflineMsg.Msg)
		}
	} else if OffLineResMsg.ResCode == 400 {
		fmt.Println("沒有離線留言")
	} else {
		fmt.Println("獲取留言未知錯誤", err)
	}

}
func (thisF *SmsProcess) GSendMessage(msg string) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("發送訊息錯誤", err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var GSmsMsg message.GSmsMsg = message.GSmsMsg{UserId: MyId, Msg: msg}
	data, err := json.Marshal(GSmsMsg)
	if err != nil {
		fmt.Println("發送訊息序列化失敗", err)
		return
	}
	tf.WritePkg(data, message.GSmsMsgType)

	var GSmsResMsg message.GSmsResMsg
	MsgData, err := tf.ReadPkg()
	json.Unmarshal([]byte(MsgData.Data), &GSmsResMsg)
	if GSmsResMsg.ResCode == 200 {
		fmt.Println("傳送成功!")
	} else if GSmsResMsg.ResCode == 400 {
		fmt.Println(GSmsResMsg.Error)
	} else {
		fmt.Println("傳送訊息未知的錯誤", err)
	}
}
