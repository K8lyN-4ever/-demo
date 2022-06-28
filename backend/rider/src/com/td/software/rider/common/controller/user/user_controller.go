package user_controller

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"gopkg.in/gin-gonic/gin.v1"
	"rider/src/com/td/software/rider/common/mapper"
	result "rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/service"
	"rider/src/com/td/software/rider/common/util/order"
	riderServices "rider/src/com/td/software/rider/rider/service"
	"strconv"
)

var store = sessions.NewCookieStore([]byte("token"))
var riderSer = riderServices.NewRiderServiceInstance()

func Register(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	typeStr := c.Query("type")
	user := mapper.User{}
	user.Account = account
	user.Password = password
	if typeStr != "" {
		loginType, err := strconv.Atoi(c.Query("type"))
		if err != nil {
			return
		}
		user.Type = loginType
	}
	data := service.UserService().Register(&user)
	c.JSON(200, data)
}

func Login(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		return
	}
	session.Options.MaxAge = 0
	if session.Values["flag"] != nil {
		c.JSON(200, result.GetFail("请勿重复登录"))
		return
	}
	account := c.Query("account")
	password := c.Query("password")
	longitudeStr := c.Query("longitude")
	latitudeStr := c.Query("latitude")

	data := service.UserService().Login(&mapper.User{
		Account:  account,
		Password: password,
	})

	bytes, err := json.Marshal(data.Data)
	info := service.LoginInfo{}
	err = json.Unmarshal(bytes, &info)

	if data.Code == "0" {
		if info.Flag != 1 {
			if longitudeStr != "" && latitudeStr != "" {
				var longitude float64
				var latitude float64
				var err error
				longitude, err = strconv.ParseFloat(longitudeStr, 64)
				latitude, err = strconv.ParseFloat(latitudeStr, 64)
				if err != nil {
					c.JSON(500, result.GetFail("参数格式错误"))
				}
				riderSer.SetLocation(account, order.Place{
					Longitude: longitude,
					Latitude:  latitude,
				})
				riderSer.SetFlag(account, "false")
			}
		}
	}
	if err != nil {
		return
	}
	session.Values["flag"] = info.Flag
	session.Values["account"] = account
	err = session.Save(c.Request, c.Writer)
	c.JSON(200, data)
}

func Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		return
	}
	if session.Values["flag"] != nil {
		session.Options.MaxAge = -1
		err := session.Save(c.Request, c.Writer)
		if err != nil {
			return
		}
		c.JSON(200, result.GetSuccess("注销成功"))
		return
	}
	c.JSON(200, result.GetFail("请先登录"))
}
