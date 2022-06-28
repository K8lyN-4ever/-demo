package mapper

import (
	"rider/src/com/td/software/rider/common/resources"
	"rider/src/com/td/software/rider/common/util"
	"sync"
)

const tableName = "riders.rider"

type User struct {
	id       int64  `gorm:"column:id"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	Type     int    `gorm:"column:type"`
}

func (User) TableName() string {
	return tableName
}

type UserMapper struct {
}

var riderMapper *UserMapper
var riderOnce sync.Once

func NewUserMapperInstance() *UserMapper {
	riderOnce.Do(
		func() {
			riderMapper = &UserMapper{}
		})
	return riderMapper
}

func (*UserMapper) QueryUserByAccount(account string) (*User, error) {
	var rider User
	err := resources.Database.Where("account = ?", account).Find(&rider).Error
	if err != nil {
		util.Logger.Error("find rider by account err:" + err.Error())
		return nil, err
	}
	return &rider, nil
}

func (*UserMapper) QueryAdminsAccount() ([]string, error) {
	var users []User
	err := resources.Database.Select("account").Where("type=  1").Find(&users).Error
	accounts := make([]string, len(users))
	for i, user := range users {
		accounts[i] = user.Account
	}
	if err != nil {
		//util.Logger.Error("find admins err:" + err.Error())
		return nil, err
	}
	return accounts, nil
}

func (*UserMapper) QueryUsers() ([]User, error) {
	var riders []User
	err := resources.Database.Find(&riders).Error
	if err != nil {
		util.Logger.Error("find riders err:" + err.Error())
		return nil, err
	}
	return riders, nil
}

func (*UserMapper) CreateUser(rider *User) error {
	if err := resources.Database.Create(rider).Error; err != nil {
		util.Logger.Error("create rider err:" + err.Error())
		return err
	}
	return nil
}

func (*UserMapper) DestroyUser(account string) error {
	if err := resources.Database.Where("account = ?", account).Delete(&User{}).Error; err != nil {
		util.Logger.Error("destroy rider err:" + err.Error())
		return err
	}
	return nil
}

func (*UserMapper) SetType(user *User) error {
	if err := resources.Database.Model(user).Where("account = ?", user.Account).Update("type", user.Type).Error; err != nil {
		util.Logger.Error("set user type err:" + err.Error())
		return err
	}
	return nil
}

func (*UserMapper) SetUser(rider *User) (*User, error) {
	if err := resources.Database.Model(rider).Where("account = ?", rider.Account).Updates(rider).Error; err != nil {
		util.Logger.Error("set rider err:" + err.Error())
		return nil, err
	}
	return rider, nil
}
