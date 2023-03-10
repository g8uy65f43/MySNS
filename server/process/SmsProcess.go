package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
	"pro05/server/model"
)

type SmsProcess struct {
	Conn net.Conn
}

func (thisF *SmsProcess) SendMessage(msgData message.MessageData) {
	var SmsMsg message.SmsMsg
	var SmsResMsh message.SmsResMsg
	var tf utils.Transfer
	json.Unmarshal([]byte(msgData.Data), &SmsMsg)
	err := Rdb.SearchUserById(SmsMsg.OtherId)
	if err == model.ERROR_USER_NOTEXISTS {
		fmt.Println(err)
		SmsResMsh.ResCode = 400
		SmsResMsh.Error = err.Error()
	} else if SmsMsg.UserId == SmsMsg.OtherId {
		SmsResMsh.ResCode = 300
		SmsResMsh.Error = "你要跟自己對話幹嘛?"
	} else {
		getUser, err := Rdb.getUserById(SmsMsg.UserId)
		if err != nil {
			fmt.Println(err)
			return
		}
		SmsMsg.UserName = getUser.UserName
		data, err := json.Marshal(SmsMsg)
		if err != nil {
			fmt.Println("轉送訊息序列化失敗")
		}
		var onLineFlag bool
		for i, v := range AllOnlineInfo {
			if i == SmsMsg.OtherId {
				tf.Conn = v.Conn
				tf.WritePkg([]byte(data), message.SmsMsgType)
				onLineFlag = true
				SmsResMsh.ResCode = 200
				break
			}
		}
		if !onLineFlag {
			MsgRdb.saveMsg(SmsMsg.OtherId, string(data))
			SmsResMsh.ResCode = 220
		}
	}
	tf.Conn = thisF.Conn
	ResData, err := json.Marshal(SmsResMsh)
	if err != nil {
		fmt.Println(err)
	}
	tf.WritePkg(ResData, message.SmsResMsgType)
}
func (thisF *SmsProcess) GetOffLineMsg(msgData message.MessageData) {
	var OffLineMsg message.OffLineMsg
	var OffLineResMsg message.OffLineResMsg
	var tf utils.Transfer = utils.Transfer{Conn: thisF.Conn}
	json.Unmarshal([]byte(msgData.Data), &OffLineMsg)
	data := MsgRdb.getOffLineMsg(OffLineMsg.UserId)
	OffLineResMsg.Msgs = data
	if len(data) == 0 {
		OffLineResMsg.ResCode = 400
	} else {
		OffLineResMsg.ResCode = 200
	}

	ResMsg, err := json.Marshal(OffLineResMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	tf.WritePkg([]byte(ResMsg), message.OffLineResMsgType)
}
func (thisF *SmsProcess) GSendMessage(msgData message.MessageData) {
	var GSmsMsg message.GSmsMsg
	var GSmsResMsg message.GSmsResMsg

	var tf utils.Transfer
	json.Unmarshal([]byte(msgData.Data), &GSmsMsg)
	data := Rdb.getAllUser(GSmsMsg.UserId)
	getUser, err := Rdb.getUserById(GSmsMsg.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	GSmsMsg.UserName = getUser.UserName
	GMsgToGroup, err := json.Marshal(GSmsMsg)
	if err != nil {
		fmt.Println("轉送訊息序列化失敗")
	}
	for _, v := range data {
		onLineData, ok := AllOnlineInfo[v]
		if ok {
			tf.Conn = onLineData.Conn
			tf.WritePkg([]byte(GMsgToGroup), message.GSmsMsgType)
		} else {
			MsgRdb.saveMsg(v, string(GMsgToGroup))
		}
		if err != nil {
			GSmsResMsg.ResCode = 400
		} else {
			GSmsResMsg.ResCode = 200
		}
		tf.Conn = thisF.Conn
		ResData, err := json.Marshal(GSmsResMsg)
		if err != nil {
			fmt.Println(err)
		}
		tf.WritePkg(ResData, message.GSmsResMsgType)
	}
}
