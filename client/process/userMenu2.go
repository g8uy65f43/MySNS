package process

import (
	"fmt"
	"os"
	"strconv"
)

func menu2() {
	var key int
	for {
		fmt.Println("-------------------------歡迎登陸多人聊天系統-------------------------")
		fmt.Println("\t\t\t 1 : 顯示所有在線列表 ")
		fmt.Println("\t\t\t 2 : 顯示好友在線列表 ")
		fmt.Println("\t\t\t 3 : 發送訊息 ")
		fmt.Println("\t\t\t 4 : 新增好友 ")
		fmt.Println("\t\t\t 5 : 退出系統 ")
		fmt.Println("\t\t\t  請選擇(1-5): ")
		fmt.Printf("\n\n")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("顯示所有在線列表")
			fmt.Println("當前在線的有:")
			for i, v := range AllOnlineInfo {
				fmt.Printf("會員ID:\t %v\t| 會員姓名: \t%v\t\n", i, v.UserName)
			}

		case 2:
			fmt.Println("顯示好友在線列表:")
			for i, v := range FriendList {
				index, _ := strconv.ParseInt(i, 10, 64)
				if AllOnlineInfo[int(index)].UserStatus == 1 {
					fmt.Printf("會員ID:\t %v\t| 會員姓名: \t%v\t|狀態:\t%v\t\n", i, v, "在線")
				} else {
					fmt.Printf("會員ID:\t %v\t| 會員姓名: \t%v\t|狀態:\t%v\t\n", i, v, "離線")
				}
			}
		case 3:
			fmt.Println("發送訊息")
			menu3()
		case 4:
			fmt.Println("新增好友")
			fmt.Println("請輸入要添加的好友ID")
			var Fid int
			fmt.Scanln(&Fid)
			up2 := UserProcess{}
			up2.AddFriend(Fid)
		case 5:
			fmt.Println("退出系統")
			LogOut(MyId)
			os.Exit(0)
		default:
			fmt.Println("你的輸入有誤，請重新輸入。")
			continue
		}
	}
}
