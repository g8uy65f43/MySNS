package message

const (
	LoginMsgType        = "LoginMsg"
	LoginResMsgType     = "LoginResMsg"
	RegisterMsgType     = "RegisterMsg"
	RegisterResMsgType  = "RegisterResMsg"
	NotifyMsgType       = "NotifyMsg"
	NotifyResMsgType    = "NotifyResMsg"
	AddFriendMsgType    = "AddFriendMsg"
	AddFriendMsgResType = "AddFriendResMsg"
	LogOutMsgType       = "LogOutMsg"
	SmsMsgType          = "SmsMsg"
	SmsResMsgType       = "SmsResMsg"
	GSmsMsgType         = "GSmsMsg"
	GSmsResMsgType      = "GSmsResMsg"
	OffLineMsgType      = "OffLineMsg"
	OffLineResMsgType   = "OffLineResMsg"
)
const (
	User_Offline = iota
	User_Online
	UserBusyStatus
)
const (
	User_Gender_Male = iota
	User_Gender_Female
)

// -------父類模型
type Userinfo struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserGender int    `json:"userGender"`
}
type UserModel struct {
	Userinfo
	UserStatus int `json:"userStatus"`
}
type ResModel struct {
	ResCode int    `json:"resCode"`
	Error   string `json:"error"`
}

// ---------主要結構體
type OnlineInfoDemo struct {
	UserId     int    `json:"userId"`
	UserStatus int    `json:"userStatus"`
	UserName   string `json:"userName"`
}
type Message struct {
	Bytes   uint32 `json:"bytes"`
	MsgData string `json:"msgData"`
}

type MessageData struct {
	MsgType string `json:"msgType"`
	Data    string `json:"data"`
}

type LoginMsg struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type LoginResMsg struct {
	ResModel
	AllOnlineId map[int]OnlineInfoDemo `json:"allOnlineId"`
	FriendList  map[string]string      `json:"friendList"`
	UserName    string                 `json:"userName"`
}

type RegisterMes struct {
	UserModel
}
type RegisterResMes struct {
	ResModel
}
type NotifyMsg struct {
	Userinfo
}
type NotifyResMsg struct {
	Userinfo
}
type AddFriendMsg struct {
	UserId   int `json:"userId"`
	FriendId int `json:"FriendId"`
}
type AddFriendResMsg struct {
	ResModel
	NewFriend map[int]Userinfo `json:"newFriend"`
}
type LogOutMsg struct {
	UserId int `json:"userId"`
}
type SmsMsg struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	OtherId  int    `json:"otherId"`
	Msg      string `json:"msg"`
}
type OffLineMsg struct {
	UserId int `json:"userId"`
}
type OffLineResMsg struct {
	ResModel
	Msgs []string `json:"userId"`
}
type GSmsMsg struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	Msg      string `json:"msg"`
}
type SmsResMsg struct {
	ResModel
}
type GSmsResMsg struct {
	ResModel
}
