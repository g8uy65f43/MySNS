package main

import (
	"fmt"
	"net"
	"pro05/server/process"
)

// 服務客戶端協成
func processGorouTine(conn net.Conn) {
	//延遲關閉conn
	defer conn.Close()
	processor := &Process{Conn: conn}
	for {
		err := processor.process()
		if err != nil {
			fmt.Println(process.AllOnlineInfo)
			fmt.Println("客戶端退出，服務端也退出...")
			return
		}
	}
}

func main() {
	fmt.Println("服務器在8889端口監聽...")
	//綁定監聽端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen Err = ", err)
		return
	}
	//延遲關端口
	defer listen.Close()
	for {
		fmt.Println("等待客戶端來連接服務器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("等待客戶連接出錯 Err =", err)
		}
		go processGorouTine(conn)
	}

}
