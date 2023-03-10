package process

import (
	"context"
	"encoding/json"
	"fmt"
	"pro05/commom/message"
	"pro05/server/model"
	"strconv"

	"github.com/go-redis/redis/v9"
)

var Ctx = context.Background()
var Rdb = CreateRdb()
var FriendRdb = CreateFriendRdb()
var MsgRdb = CreateMsgRdb()

type UserDao struct {
	Rdb *redis.Client
}

func CreateRdb() *UserDao {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "g8uy65f43", // no password set
		DB:       0,           // use default DB
	})
	var userDao UserDao = UserDao{Rdb: rdb}
	return &userDao
}

func CreateFriendRdb() *UserDao {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "g8uy65f43", // no password set
		DB:       1,           // use default DB
	})
	var userDao UserDao = UserDao{Rdb: rdb}
	return &userDao
}
func CreateMsgRdb() *UserDao {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "g8uy65f43", // no password set
		DB:       2,           // use default DB
	})
	var userDao UserDao = UserDao{Rdb: rdb}
	return &userDao
}
func (thisF *UserDao) getUserById(UserId int) (user *message.Userinfo, err error) {
	res, err := thisF.Rdb.HGet(Ctx, "users", fmt.Sprint(UserId)).Result()
	if err != nil {
		fmt.Println("沒有查詢的用戶")
		err = model.ERROR_USER_NOTEXISTS
		return
	}
	user = &message.Userinfo{}
	json.Unmarshal([]byte(res), &user)
	return
}
func (thisF *UserDao) Login(UserId int, UserPwd string) (user *message.Userinfo, err error) {
	user, err = thisF.getUserById(UserId)
	if err != nil {
		return
	}
	if UserPwd != user.UserPwd {
		err = model.ERROR_USER_PWD
		return
	}
	return
}
func (thisF *UserDao) SearchUserById(UserId int) (err error) {
	ok := thisF.Rdb.HExists(Ctx, "users", fmt.Sprint(UserId)).Val()
	if !ok {
		err = model.ERROR_USER_NOTEXISTS
		return
	}
	return
}
func (thisF *UserDao) Register(UserId, UserGender int, UserPwd, UserName string) (user *message.Userinfo, err error) {
	ok := thisF.Rdb.HExists(Ctx, "users", fmt.Sprint(UserId)).Val()
	if ok {
		err = model.ERROR_USER_EXISTS
		return
	}
	user = &message.Userinfo{UserId: UserId, UserPwd: UserPwd, UserName: UserName, UserGender: UserGender}

	registerData, err := json.Marshal(&user)
	thisF.Rdb.HSet(Ctx, "users", UserId, string(registerData))
	return
}
func (thisF *UserDao) Add(UserId, FriendId int) (user *message.Userinfo, err error) {
	allUserData := CreateRdb()

	data, err := allUserData.getUserById(FriendId)
	if err != nil {
		err = model.ERROR_USER_NOTEXISTS
		return
	}
	ok := thisF.Rdb.HExists(Ctx, "users", fmt.Sprint(UserId)).Val()
	if ok {
		err = model.ERROR_FRIEND_EXISTS
		return
	}
	user = &message.Userinfo{}
	user.UserName = data.UserName
	user.UserId = data.UserId
	thisF.Rdb.HSet(Ctx, fmt.Sprint(UserId), data.UserId, data.UserName)
	return
}
func (thisF *UserDao) getFriendList(UserId int) (user map[string]string) {
	data := thisF.Rdb.HGetAll(Ctx, fmt.Sprint(UserId)).Val()
	user = make(map[string]string)
	for i, v := range data {
		user[i] = v
	}
	return
}
func (thisF *UserDao) saveMsg(UserId int, msg string) {
	thisF.Rdb.RPush(Ctx, fmt.Sprint(UserId), msg)
}
func (thisF *UserDao) getOffLineMsg(UserId int) (msgs []string) {
	ok := thisF.Rdb.LRange(Ctx, fmt.Sprint(UserId), 0, -1).Val()
	msgs = ok
	thisF.Rdb.Del(Ctx, fmt.Sprint(UserId))
	return
}
func (thisF *UserDao) getAllUser(UserId int) (userIDArr []int) {
	ok := thisF.Rdb.HGetAll(Ctx, "users").Val()
	for i, _ := range ok {
		data, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			fmt.Println("獲取全部用戶，數字轉換錯誤")
			return
		}
		if int(data) == UserId {
			continue
		} else {
			userIDArr = append(userIDArr, int(data))
		}
	}
	return userIDArr
}
