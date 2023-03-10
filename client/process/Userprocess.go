package process

import (
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
	"pro05/commom/utils"
)

type UserProcess struct {
}

var MyId int

func (thisF *UserProcess) Login(UserId int, UserPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Login(C)與伺服器連接失敗")
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var LoginMsg = message.LoginMsg{UserId: UserId, UserPwd: UserPwd}
	data, err := json.Marshal(LoginMsg)
	if err != nil {
		fmt.Println("LoginMsg序列化失敗", err)
	}
	tf.WritePkg(data, message.LoginMsgType)
	msg, err := tf.ReadPkg()
	if err != nil {
		return
	}
	var resMesData message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &resMesData)
	if err != nil {
		fmt.Println(err)
		return
	}
	var userName = resMesData.UserName
	if resMesData.ResCode == 200 {
		fmt.Println("登入成功!")
		fmt.Println("歡迎回來!", userName)
		MyId = UserId
		for i, v := range resMesData.AllOnlineId {
			AllOnlineInfo[i] = message.OnlineInfoDemo{UserId: v.UserId, UserStatus: v.UserStatus, UserName: v.UserName}
		}
		FriendList = resMesData.FriendList
		go server(conn)
		menu2()
	} else if resMesData.ResCode == 300 {
		fmt.Println("密碼錯誤!")
	} else if resMesData.ResCode == 400 {
		fmt.Println("沒有該用戶!")
	} else {
		fmt.Println("遇到未知的錯誤")
	}

	return
}
func (thisF *UserProcess) Register(UserId, UserGender int, UserPwd, UserName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Login(C)與伺服器連接失敗")
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var registerMsg message.RegisterMes
	registerMsg.UserGender = UserGender
	registerMsg.UserId = UserId
	registerMsg.UserName = UserName
	registerMsg.UserPwd = UserPwd
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("registerMsg序列化失敗", err)
	}
	tf.WritePkg(data, message.RegisterMsgType)
	msg, err := tf.ReadPkg()
	if err != nil {
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(msg.Data), &registerResMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	if registerResMes.ResCode == 200 {
		fmt.Println("註冊成功!請重新登入")
	} else if registerResMes.ResCode == 300 {
		fmt.Println("輸入格式錯誤!")
	} else if registerResMes.ResCode == 400 {
		fmt.Println("該用戶已存在!")
	} else {
		fmt.Println("遇到未知的錯誤")
	}

	return
}
func LogOut(UserId int) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	var LogOutMsg message.LogOutMsg = message.LogOutMsg{UserId: UserId}
	data, err := json.Marshal(LogOutMsg)
	if err != nil {
		fmt.Println("登出回傳失敗")
	}
	tf.WritePkg([]byte(data), message.LogOutMsgType)
	return
}
