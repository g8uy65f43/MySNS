package main

import (
	"fmt"
	"os"
	"pro05/client/process"
)

var userId, userGender int
var userPwd, userName string

func main() {
	var key int
	for {
		fmt.Println("-------------------------歡迎登陸多人聊天系統-------------------------")
		fmt.Println("\t\t\t 1 : 登入聊天系統 ")
		fmt.Println("\t\t\t 2 : 註冊用戶 ")
		fmt.Println("\t\t\t 3 : 登出系統 ")
		fmt.Println("\t\t\t  請選擇(1-3): ")
		fmt.Printf("\n\n")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			fmt.Println("請輸入用戶ID號:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("請輸入用戶密碼:")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			key = 0
		case 2:
			fmt.Println("註冊用戶")
			fmt.Println("請輸入用戶ID號:")
			fmt.Scanln(&userId)
			fmt.Println("請輸入用戶密碼:")
			fmt.Scanln(&userPwd)
			fmt.Println("請輸入用戶性別:(男姓:0,女性:1)")
			fmt.Scanln(&userGender)
			if userGender < 0 || userGender > 1 {
				fmt.Println("錯誤的輸入")
				return
			}
			fmt.Println("請輸入用戶名稱:")
			fmt.Scanln(&userName)
			up := &process.UserProcess{}
			up.Register(userId, userGender, userPwd, userName)
			os.Exit(0)
		case 3:
			fmt.Println("退出系統")
			os.Exit(0)
		default:
			fmt.Println("你的輸入有誤，請重新輸入。")
			continue
		}
	}
}
