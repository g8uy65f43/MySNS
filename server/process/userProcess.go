package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
	"pro05/server/model"
)

type UserProcess struct {
	Conn net.Conn
}

func (thisF *UserProcess) ServerProcessLogin(msgData message.MessageData) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msgData.Data), &loginMsg)
	if err != nil {
		fmt.Println(err)
	}

	var loginResMsg message.LoginResMsg = message.LoginResMsg{}
	userData, err := Rdb.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		loginResMsg.Error = err.Error()
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.ResCode = 400
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.ResCode = 300
		}
		fmt.Println("登入失敗")
	} else {
		fmt.Println("登入成功")
		AllOnlineInfo[userData.UserId] = ServerOnlineInfo{UserId: userData.UserId, Conn: thisF.Conn, UserStatus: 1, UserName: userData.UserName}
		fmt.Printf("id為:%v,姓名為%v,密碼為:%v\n", userData.UserId, userData.UserName, userData.UserPwd)
		loginResMsg.UserName = userData.UserName
		loginResMsg.ResCode = 200
		loginResMsg.AllOnlineId = make(map[int]message.OnlineInfoDemo)
		loginResMsg.FriendList = make(map[string]string)
		var FriendList map[string]string = make(map[string]string)
		FriendList = FriendRdb.getFriendList(loginMsg.UserId)
		loginResMsg.FriendList = FriendList
		for id, v := range AllOnlineInfo {
			if id == userData.UserId || v.UserStatus != 1 {
				continue
			}
			loginResMsg.AllOnlineId[id] = message.OnlineInfoDemo{UserId: v.UserId, UserStatus: v.UserStatus, UserName: v.UserName}
		}
		NotifyOthersOnlineUser(userData.UserId)
	}
	tf := utils.Transfer{
		Conn: thisF.Conn,
	}
	data, err := json.Marshal(loginResMsg)
	tf.WritePkg(data, message.LoginResMsgType)
	return
}

func (thisF *UserProcess) ServerProcessRegister(msgData message.MessageData) (err error) {
	var registerMsg message.RegisterMes
	err = json.Unmarshal([]byte(msgData.Data), &registerMsg)
	if err != nil {
		fmt.Println(err)
	}

	var registerResMsg message.RegisterResMes = message.RegisterResMes{}
	userData, err := Rdb.Register(registerMsg.UserId, registerMsg.UserGender, registerMsg.UserPwd, registerMsg.UserName)
	if err != nil {
		registerResMsg.Error = err.Error()
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.ResCode = 400
		} else if err == model.ERROR_USER_STYLE_ERROR {
			registerResMsg.ResCode = 300
		}
		fmt.Println("註冊失敗")
	} else {
		fmt.Println("註冊成功")
		registerResMsg.ResCode = 200
		fmt.Println(userData)
	}
	tf := utils.Transfer{
		Conn: thisF.Conn,
	}
	data, err := json.Marshal(registerResMsg)
	tf.WritePkg(data, message.RegisterResMsgType)
	return

}
func (thisF *UserProcess) ServerProcessAddFriend(msgData message.MessageData) (err error) {
	var AddFriend message.AddFriendMsg
	err = json.Unmarshal([]byte(msgData.Data), &AddFriend)
	if err != nil {
		fmt.Println(err)
	}
	var AddFriendResMsg message.AddFriendResMsg = message.AddFriendResMsg{}
	if AddFriend.UserId != AddFriend.FriendId {
		fromDbData, err := FriendRdb.Add(AddFriend.UserId, AddFriend.FriendId)
		if err != nil {
			AddFriendResMsg.Error = err.Error()
			if err == model.ERROR_USER_NOTEXISTS {
				AddFriendResMsg.ResCode = 400
			}
		} else {
			fmt.Println("好友添加成功")
			AddFriendResMsg.ResCode = 200
			AddFriendResMsg.NewFriend = make(map[int]message.Userinfo)
			AddFriendResMsg.NewFriend[fromDbData.UserId] = message.Userinfo{UserId: fromDbData.UserId, UserName: fromDbData.UserName}
		}
	} else {
		AddFriendResMsg.ResCode = 300
	}

	tf := utils.Transfer{
		Conn: thisF.Conn,
	}
	data, err := json.Marshal(AddFriendResMsg)
	if err != nil {
		fmt.Println(err)
	}
	tf.WritePkg(data, message.AddFriendMsgResType)
	return

}
