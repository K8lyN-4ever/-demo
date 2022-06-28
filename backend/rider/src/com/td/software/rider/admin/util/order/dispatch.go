package order

import (
	"fmt"
	sys_flag_service "rider/src/com/td/software/rider/admin/service/sys_flag"
	"rider/src/com/td/software/rider/common/service"
	orderSrv "rider/src/com/td/software/rider/common/service"
	order2 "rider/src/com/td/software/rider/common/util/order"
	ws "rider/src/com/td/software/rider/common/util/websocket"
	riderSrv "rider/src/com/td/software/rider/rider/service"
	"sync"
)

var (
	dispatchLock sync.Mutex
	riderService = riderSrv.NewRiderServiceInstance()
	orderService = orderSrv.NewOrderServiceInstance()
	userService  = service.UserService()
)

func dispatch() {
	for {
		flag := sys_flag_service.NewSysFlagServiceInstance().GetDispatchFlag()
		for flag {
			orders := orderService.GetOriROrders()
			for _, order := range orders {
				riders, err := riderService.GetOriRidersByDist(order.Src)
				if err != nil {
					continue
				}
				for _, rider := range riders {
					if !riderService.IsGrab(rider.Name) {
						continue
					}
					if res := orderService.AccOrder(order.Uuid, rider.Name); res.Code != "0" {
						continue
					}
					ids := userService.GetAdminsAccount()
					ids = append(ids, rider.Name)
					ws.WebsocketManager.Success(order2.BroadcastData{
						OrderId: order.Uuid,
						Flag:    "accept",
					}, false, ids)
					fmt.Println("订单：" + order.Uuid + "分配给了:" + rider.Name)

				}
			}
			flag = sys_flag_service.NewSysFlagServiceInstance().GetDispatchFlag()
		}
	}
}

func StopDispatch() {
	dispatchLock.Lock()
	if sys_flag_service.NewSysFlagServiceInstance().GetDispatchFlag() {
		sys_flag_service.NewSysFlagServiceInstance().SetDispatchFlag("false")
	}
	dispatchLock.Unlock()
}

func Dispatch() {
	dispatchLock.Lock()
	if !sys_flag_service.NewSysFlagServiceInstance().GetDispatchFlag() {
		sys_flag_service.NewSysFlagServiceInstance().SetDispatchFlag("true")
	}
	dispatchLock.Unlock()
}
