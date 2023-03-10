package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
)

func (thisF *UserProcess) AddFriend(UserId int) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("AddFFriend(C)與伺服器連接失敗")
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var AddFriendMsg message.AddFriendMsg = message.AddFriendMsg{UserId: MyId, FriendId: UserId}
	data, err := json.Marshal(AddFriendMsg)
	if err != nil {
		fmt.Println("好友添加失敗(C)")
		return
	}
	tf.WritePkg(data, message.AddFriendMsgType)

	var addResMsg message.AddFriendResMsg
	resData, err := tf.ReadPkg()
	json.Unmarshal([]byte(resData.Data), &addResMsg)
	if addResMsg.ResCode == 400 {
		fmt.Println(addResMsg.Error)
	} else if addResMsg.ResCode == 200 {
		fmt.Println("添加成功!")
		for i, v := range addResMsg.NewFriend {
			fmt.Println("好友:", v)
			FriendList[fmt.Sprint(i)] = v.UserName
		}
	} else if addResMsg.ResCode == 300 {
		fmt.Println("不能添加自己為好友")
	}
	return
}
