package service

import (
	"encoding/json"
	"rider/src/com/td/software/rider/common/mapper"
	"rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/util"
	"sync"
	"time"
)

type LoginInfo struct {
	Flag      int       `json:"flag"`
	LoginTime time.Time `json:"login_time"`
}

type RidersInfo struct {
	Total  int           `json:"total"`
	Riders []mapper.User `json:"riders"`
}

type UserServiceImpl struct {
}

var userService *UserServiceImpl
var userOnce sync.Once

func UserService() *UserServiceImpl {
	userOnce.Do(
		func() {
			userService = &UserServiceImpl{}
		})
	return userService
}

func NewRiderFlow(account string) *mapper.User {
	return &mapper.User{
		Account: account,
	}
}

func verification(rider *mapper.User) (bool, int, error) {
	trueRider, err := mapper.NewUserMapperInstance().QueryUserByAccount(rider.Account)
	if err != nil {
		return false, -1, err
	}
	if trueRider.Password == rider.Password {
		return true, trueRider.Type, nil
	} else {
		return false, -1, nil
	}
}

func (UserServiceImpl) Register(rider *mapper.User) *result.Result {
	if err := mapper.NewUserMapperInstance().CreateUser(rider); err != nil {
		util.Logger.Error("service=>register err:" + err.Error())
		return result.GetFail(err.Error())
	}
	return result.GetSuccess("register successful")
}

func (UserServiceImpl) Login(rider *mapper.User) *result.Result {
	ok, flag, err := verification(rider)
	if err != nil {
		return result.GetFail("登录失败")
	}
	info, err := json.Marshal(LoginInfo{
		Flag:      flag,
		LoginTime: time.Now(),
	})
	if err != nil {
		return result.GetFail("登录失败")
	}
	if ok {
		return &result.Result{
			Code: "0",
			Msg:  "登录成功",
			Data: json.RawMessage(info),
		}
	} else {
		return result.GetFail("登录失败")
	}
}

func (UserServiceImpl) GetRiders() *result.Result {
	riders, err := mapper.NewUserMapperInstance().QueryUsers()
	if err != nil {
		util.Logger.Error("service=>get riders err" + err.Error())
		return result.GetFail("获取骑手列表失败")
	}
	return &result.Result{
		Code: "0",
		Msg:  "",
		Data: RidersInfo{
			Total:  len(riders),
			Riders: riders,
		},
	}
}

func (UserServiceImpl) DeleteUser(account string) *result.Result {
	if err := mapper.NewUserMapperInstance().DestroyUser(account); err != nil {
		util.Logger.Error("service=>delete rider err" + err.Error())
		return result.GetFail("删除失败")
	}
	return result.GetSuccess("删除成功")
}

func (UserServiceImpl) SetType(user *mapper.User) *result.Result {
	if err := mapper.NewUserMapperInstance().SetType(user); err != nil {
		return result.GetFail("修改失败")
	}
	return result.GetSuccess("修改成功")
}

func (UserServiceImpl) UpdateUser(user mapper.User) *result.Result {
	if _, err := mapper.NewUserMapperInstance().SetUser(&user); err != nil {
		return result.GetFail("修改失败")
	}
	return result.GetSuccess("修改成功")
}

func (UserServiceImpl) GetAdminsAccount() []string {
	res, err := mapper.NewUserMapperInstance().QueryAdminsAccount()
	if err != nil {
		return []string{}
	}
	return res
}
