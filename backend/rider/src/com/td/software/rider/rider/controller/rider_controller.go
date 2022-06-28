package rider_controller

import (
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/gorilla/sessions"
	result "rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/service"
	rider_service "rider/src/com/td/software/rider/rider/service"
)

var store = sessions.NewCookieStore([]byte("token"))

func grab(c *gin.Context, flag string) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		c.JSON(200, result.GetSimpleFail())
		return
	}
	data := rider_service.NewRiderServiceInstance().SetFlag(session.Values["account"].(string), flag)
	c.JSON(200, data)
}

func GetPicked(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		c.JSON(200, result.GetSimpleFail())
		return
	}
	account := session.Values["account"].(string)
	data := service.NewOrderServiceInstance().GetPickedOrders(account)
	c.JSON(200, data)
}

func GetOrders(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		c.JSON(200, result.GetSimpleFail())
		return
	}
	account := session.Values["account"].(string)
	location := rider_service.NewRiderServiceInstance().GetRiderLocation(account)
	orders, _ := rider_service.NewRiderServiceInstance().GetOriOrdersByDist(location)
	res, _ := json.Marshal(orders)
	c.JSON(200, string(res))
}

func Accept(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		c.JSON(200, result.GetSimpleFail())
		return
	}
	account := session.Values["account"].(string)
	uid := c.Query("uuid")
	data := service.NewOrderServiceInstance().AccOrder(uid, account)
	c.JSON(200, data)
}

func Cancel(c *gin.Context) {
	uid := c.Query("uuid")
	data := service.NewOrderServiceInstance().CanOrder(uid)
	c.JSON(200, data)
}

func Grab(c *gin.Context) {
	grab(c, "true")
}

func UnGrab(c *gin.Context) {
	grab(c, "false")
}

func Complete(c *gin.Context) {
	uid := c.Query("uuid")
	data := service.NewOrderServiceInstance().ComOrder(uid)
	c.JSON(200, data)
}

func GetFlag(c *gin.Context) {
	session, err := store.Get(c.Request, "token")
	if err != nil {
		c.JSON(200, result.GetSimpleFail())
		return
	}
	account := session.Values["account"].(string)
	data := rider_service.NewRiderServiceInstance().IsGrab(account)
	c.JSON(200, data)
}
