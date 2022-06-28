package order

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	sys_flag_service "rider/src/com/td/software/rider/admin/service/sys_flag"
	"rider/src/com/td/software/rider/common/mapper"
	"rider/src/com/td/software/rider/common/service"
	"rider/src/com/td/software/rider/common/util/order"
	ws "rider/src/com/td/software/rider/common/util/websocket"
	riderSer "rider/src/com/td/software/rider/rider/service"
	"sync"
	"time"
)

var generateLock sync.Mutex

// 通达附近
const (
	minLongitude = 119.3660541
	maxLongitude = 119.4153504
	minLatitude  = 32.3322146
	maxLatitude  = 32.3987607
)

func getRandPlace() order.Place {
	return order.Place{
		Longitude: getRandLongitude(),
		Latitude:  getRandLatitude(),
	}
}

func getRandLongitude() float64 {
	return minLongitude + rand.Float64()*(maxLongitude-minLongitude)
}

func getRandLatitude() float64 {
	return minLatitude + rand.Float64()*(maxLatitude-minLatitude)
}

func generateRandOrder() *order.Order {
	randOrder := order.Order{
		Uuid: uuid.NewV4().String(),
		Src:  getRandPlace(),
		Tar:  getRandPlace(),
	}
	service.NewOrderServiceInstance().CreateOrder(randOrder)
	return &randOrder
}

func generate() {
	for {
		flag := sys_flag_service.NewSysFlagServiceInstance().GetGenerateFlag() && mapper.NewOrderMapperInstance().GetNum() < 10
		for flag {
			temp := generateRandOrder()
			msg, _ := json.Marshal(temp)
			geoLocation, _ := riderSer.NewRiderServiceInstance().GetOriRidersByDist(temp.Src)
			ids := make([]string, len(geoLocation))
			for i, location := range geoLocation {
				ids[i] = location.Name
			}
			admins := userService.GetAdminsAccount()
			data := order.BroadcastData{
				OrderId: string(msg),
				Flag:    "append",
			}
			ws.WebsocketManager.Success(data, false, ids)
			ws.WebsocketManager.Success(data, false, admins)
			fmt.Println("生成的订单信息：" + string(msg))
			time.Sleep(time.Duration(2) * time.Second)
			flag = sys_flag_service.NewSysFlagServiceInstance().GetGenerateFlag() && mapper.NewOrderMapperInstance().GetNum() <= 10
		}
	}
}

func Generate() {
	generateLock.Lock()
	if !sys_flag_service.NewSysFlagServiceInstance().GetGenerateFlag() {
		sys_flag_service.NewSysFlagServiceInstance().SetGenerateFlag("true")
	}
	generateLock.Unlock()
}

func StopGenerate() {
	generateLock.Lock()
	if sys_flag_service.NewSysFlagServiceInstance().GetGenerateFlag() {
		sys_flag_service.NewSysFlagServiceInstance().SetGenerateFlag("false")
	}
	generateLock.Unlock()
}
