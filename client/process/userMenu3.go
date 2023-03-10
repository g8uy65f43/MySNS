package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Sms SmsProcess

func menu3() {
	var key int

	for {
		var otherId int
		var msg string
		fmt.Println("-------------------------聊天目錄-------------------------")
		fmt.Println("\t\t\t 1 : 私訊 ")
		fmt.Println("\t\t\t 2 : 群發訊息 ")
		fmt.Println("\t\t\t 3 : 獲取離線訊息 ")
		fmt.Println("\t\t\t 4 : 返回上一頁 ")
		fmt.Println("\t\t\t  請選擇(1-4): ")
		fmt.Printf("\n\n")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("請輸入私訊對象的ID")
			fmt.Scanf("%d\n", &otherId)
			fmt.Println("請輸入要傳送的訊息")
			reader := bufio.NewReader(os.Stdin)
			msg, _ = reader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			Sms.SendMessage(otherId, msg)
		case 2:
			fmt.Println("請輸入要群發的訊息")
			reader := bufio.NewReader(os.Stdin)
			msg, _ = reader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			Sms.GSendMessage(msg)
		case 3:
			fmt.Println("獲取離線訊息")
			Sms.GetOffLineMessage()
		case 4:
			fmt.Println("返回到上一頁")
			return
		default:
			fmt.Println("你的輸入有誤，請重新輸入。")
			continue
		}
	}
}
