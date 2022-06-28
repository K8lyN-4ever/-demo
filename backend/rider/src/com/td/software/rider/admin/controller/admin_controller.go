package admin_controller

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gin-gonic/gin.v1"
	sys_flag_service "rider/src/com/td/software/rider/admin/service/sys_flag"
	orderUtil "rider/src/com/td/software/rider/admin/util/order"
	"rider/src/com/td/software/rider/common/mapper"
	result "rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/service"
	"rider/src/com/td/software/rider/common/util/order"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	srcLongitudeStr := c.Query("srcLongitude")
	srcLatitudeStr := c.Query("srcLatitude")
	tarLongitudeStr := c.Query("tarLongitude")
	tarLatitudeStr := c.Query("tarLatitude")

	var srcLongitude float64
	var srcLatitude float64
	var tarLongitude float64
	var tarLatitude float64
	var err error

	srcLongitude, err = strconv.ParseFloat(srcLongitudeStr, 64)
	srcLatitude, err = strconv.ParseFloat(srcLatitudeStr, 64)
	tarLongitude, err = strconv.ParseFloat(tarLongitudeStr, 64)
	tarLatitude, err = strconv.ParseFloat(tarLatitudeStr, 64)

	if err != nil {
		c.JSON(500, result.GetFail("请输入正确参数"))
	}

	data := service.NewOrderServiceInstance().CreateOrder(order.Order{
		Uuid: uuid.NewV4().String(),
		Src: order.Place{
			Longitude: srcLongitude,
			Latitude:  srcLatitude,
		},
		Tar: order.Place{
			Longitude: tarLongitude,
			Latitude:  tarLatitude,
		},
	})
	c.JSON(200, data)
}

func DeleteOrder(c *gin.Context) {
	uid := c.Query("uuid")
	data := service.NewOrderServiceInstance().DelROrders(uid)
	c.JSON(200, data)
}

func GenerateOrder(c *gin.Context) {
	orderUtil.Generate()
	c.JSON(200, result.GetSuccess("开始生成"))
}

func StopGenerate(c *gin.Context) {
	orderUtil.StopGenerate()
	c.JSON(200, result.GetSuccess("停止生成"))
}

func GenerateFlag(c *gin.Context) {
	flag := sys_flag_service.NewSysFlagServiceInstance().GetGenerateFlag()
	c.JSON(200, flag)
}

func StartDispatch(c *gin.Context) {
	orderUtil.Dispatch()
	c.JSON(200, result.GetSuccess("开启自动分配"))
}

func StopDispatch(c *gin.Context) {
	orderUtil.StopDispatch()
	c.JSON(200, result.GetSuccess("停止自动分配"))
}

func DispatchFlag(c *gin.Context) {
	flag := sys_flag_service.NewSysFlagServiceInstance().GetDispatchFlag()
	c.JSON(200, flag)
}

func GetOrders(c *gin.Context) {
	data := service.NewOrderServiceInstance().GetOrders()
	c.JSON(200, data)
}

func GetRiders(c *gin.Context) {
	data := service.UserService().GetRiders()
	c.JSON(200, data)
}

func DeleteRider(c *gin.Context) {
	account := c.Query("account")
	data := service.UserService().DeleteUser(account)
	c.JSON(200, data)
}

func UpdatePassword(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	data := service.UserService().UpdateUser(mapper.User{
		Account:  account,
		Password: password,
		Type:     0,
	})
	c.JSON(200, data)
}

func SetUserType(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	t := c.Query("type")
	ty, err := strconv.Atoi(t)
	if err != nil {
		c.JSON(500, result.GetFail("参数类型不正确"))
	}
	user := mapper.User{
		Account:  account,
		Password: password,
		Type:     ty,
	}
	data := service.UserService().SetType(&user)
	c.JSON(200, data)
}
