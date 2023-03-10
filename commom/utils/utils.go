package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"pro05/commom/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (thisF *Transfer) WritePkg(data []byte, Type string) (err error) {
	//計算
	var MsgData = message.MessageData{MsgType: Type, Data: string(data)}
	MarshalData, err := json.Marshal(MsgData)
	if err != nil {
		fmt.Println("Login-MsgData序列化失敗", err)
	}
	var MsgSize uint32 = uint32(len(MarshalData))
	binary.BigEndian.PutUint32(thisF.Buf[:4], MsgSize)
	var Msg = message.Message{Bytes: MsgSize, MsgData: string(MarshalData)}
	data, err = json.Marshal(Msg)
	if err != nil {
		fmt.Println("Login-MsgData序列化失敗", err)
	}
	thisF.Conn.Write(data)
	return
}

func (thisF *Transfer) ReadPkg() (msgData message.MessageData, err error) {
	n, err := thisF.Conn.Read(thisF.Buf[:])
	if err != nil {
		return
	}
	var Msg message.Message
	json.Unmarshal(thisF.Buf[:n], &Msg)
	if int(Msg.Bytes) != len(Msg.MsgData) {
		fmt.Println("唉呀!丟包了")
	}
	json.Unmarshal([]byte(Msg.MsgData), &msgData)
	return
}
