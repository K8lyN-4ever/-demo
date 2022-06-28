package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	// 导入session存储引擎
	"net/http"
	admin_controller "rider/src/com/td/software/rider/admin/controller"
	"rider/src/com/td/software/rider/admin/util/order"
	user_controller "rider/src/com/td/software/rider/common/controller/user"
	ws_controller "rider/src/com/td/software/rider/common/controller/ws"
	ws "rider/src/com/td/software/rider/common/util/websocket"
	rider_controller "rider/src/com/td/software/rider/rider/controller"
)

const baseUrl = "/api"
const wsUrl = baseUrl + "/ws"
const userUrl = baseUrl + "/user"
const adminUrl = baseUrl + "/admin"
const riderUrl = baseUrl + "/rider"

// 跨域中间件
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Router() error {

	var err error

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors())

	webSocket := r.Group(wsUrl)
	{
		webSocket.GET("/connect", ws_controller.Client)
	}

	user := r.Group(userUrl)
	{
		user.GET("/register", user_controller.Register)
		user.GET("/login", user_controller.Login)
		user.GET("/logout", user_controller.Logout)
	}

	admin := r.Group(adminUrl)
	{
		admin.GET("/getUsers", admin_controller.GetRiders)
		admin.GET("/getOrders", admin_controller.GetOrders)
		admin.GET("/deleteUser", admin_controller.DeleteRider)
		admin.GET("/updateUser", admin_controller.SetUserType)
		admin.GET("/updatePassword", admin_controller.UpdatePassword)
		admin.GET("/createOrder", admin_controller.CreateOrder)
		admin.GET("/deleteOrder", admin_controller.DeleteOrder)
		admin.GET("/generateFlag", admin_controller.GenerateFlag)
		admin.GET("/startGenerate", admin_controller.GenerateOrder)
		admin.GET("/stopGenerate", admin_controller.StopGenerate)
		admin.GET("/dispatchFlag", admin_controller.DispatchFlag)
		admin.GET("/startDispatch", admin_controller.StartDispatch)
		admin.GET("/stopDispatch", admin_controller.StopDispatch)
	}

	rider := r.Group(riderUrl)
	{
		rider.GET("/getPicked", rider_controller.GetPicked)
		rider.GET("/getOrders", rider_controller.GetOrders)
		rider.GET("/accept", rider_controller.Accept)
		rider.GET("/cancel", rider_controller.Cancel)
		rider.GET("/complete", rider_controller.Complete)
		rider.GET("/grab", rider_controller.Grab)
		rider.GET("/unGrab", rider_controller.UnGrab)
		rider.GET("/getFlag", rider_controller.GetFlag)
	}

	order.Start()
	go ws.WebsocketManager.Start()

	if err != nil {
		return err
	}

	err = r.Run()
	if err != nil {
		return err
	}
	return nil
}
