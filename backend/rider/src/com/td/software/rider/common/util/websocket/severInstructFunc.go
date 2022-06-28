/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-22 11:42:19
 * @LastEditors: shahao
 * @LastEditTime: 2021-04-27 13:53:18
 */
package websocket

import (
	"encoding/json"
	"fmt"
	"rider/src/com/td/software/rider/common/util/order"
	"rider/src/com/td/software/rider/rider/service"
)

type ServerMethod struct {
}

//设备状态
// func (m *ServerMethod) EquipmentStatus(params ReadData) {
// 	WebsocketManager.Success(params.Actioncode, 21, params.IsBroadCast, params.ClientIDs)
// }

var riderService = rider_service.NewRiderServiceInstance()

//心跳包
func (m *ServerMethod) HeartBeat(params ReadData) {
	WebsocketManager.Success(true, params.IsBroadCast, params.ClientIDs)
}

func (m *ServerMethod) Demo(params ReadData) {
	str, _ := json.Marshal(params)
	fmt.Println(str)
	fmt.Println("demo")
}

func (m *ServerMethod) UpdateLocation(params ReadData) {
	var location order.Place
	if err := json.Unmarshal([]byte(params.Location), &location); err != nil {
		fmt.Printf("location 数据解析失败\n")
	}
	res := riderService.SetLocation(params.ClientIDs[0], location)
	WebsocketManager.Success(res, false, params.ClientIDs)
}

func (m *ServerMethod) Logout(params ReadData) {
	WebsocketManager.UnRegister <- &Client{
		Id:      params.ClientIDs[0],
		Socket:  nil,
		Message: nil,
	}
}
