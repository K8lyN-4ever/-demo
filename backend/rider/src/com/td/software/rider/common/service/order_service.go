package service

import (
	"encoding/json"
	"fmt"
	"rider/src/com/td/software/rider/common/mapper"
	result "rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/util/order"
	ws "rider/src/com/td/software/rider/common/util/websocket"
	"rider/src/com/td/software/rider/rider/service"
	"sync"
)

var orderMapper = mapper.NewOrderMapperInstance()
var riderSer = rider_service.NewRiderServiceInstance()

type OrderService struct {
}

var orderService *OrderService
var orderOnce sync.Once
var wg = sync.WaitGroup{}

func NewOrderServiceInstance() *OrderService {
	orderOnce.Do(
		func() {
			orderService = &OrderService{}
		})
	return orderService
}

func delROrders(uuid string) bool {
	wg.Add(4)
	ch := make(chan string)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.DelOrder(uuid)
		if err != nil {
			fmt.Println(err.Error())
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.DelDestLocation(uuid)
		if err != nil {
			fmt.Println(err.Error())
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.DelShopLocation(uuid)
		if err != nil {
			fmt.Println(err.Error())
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.DelOrderFlag(uuid)
		if err != nil {
			fmt.Println(err.Error())
			c <- err.Error()
		}
	}(ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	if len(ch) != 0 {
		return true // 出现err返回true
	}
	return false // 否则返回false
}

func broadcast(uuid string, data order.BroadcastData) {
	location := riderSer.GetOrderLocation(uuid)
	dist, err := riderSer.GetOriRidersByDist(location)
	if err != nil {
		return
	}
	ids := make([]string, len(dist))
	for i, d := range dist {
		ids[i] = d.Name
	}
	admins := userService.GetAdminsAccount()
	ws.WebsocketManager.Success(data, false, ids)
	ws.WebsocketManager.Success(data, false, admins)
}

func (service OrderService) CreateOrder(order order.Order) *result.Result {
	wg.Add(4)
	ch := make(chan string)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.CreateOrder(order)
		if err != nil {
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.AddShopLocation(order.Uuid, order.Src)
		if err != nil {
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.AddDestLocation(order.Uuid, order.Tar)
		if err != nil {
			c <- err.Error()
		}
	}(ch)
	go func(c chan string) {
		defer wg.Done()
		err := orderMapper.AddOrderFlag(order.Uuid)
		if err != nil {
			c <- err.Error()
		}
	}(ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	if len(ch) != 0 {
		return result.GetFail("创建失败")
	}
	return &result.Result{
		Code: "0",
		Msg:  "创建成功",
		Data: order,
	}
}

func (service OrderService) GetOrder(uuid string) *result.Result {
	res, err := orderMapper.GetOrder(uuid)
	data, err := json.Marshal(res)
	if err != nil {
		return result.GetFail("查询失败")
	}
	return &result.Result{
		Code: "0",
		Msg:  "查询成功",
		Data: data,
	}
}

func (service OrderService) GetOrders() result.Result {
	res, err := orderMapper.GetOrders()
	data, _ := json.Marshal(res)
	if err != nil {
		return *result.GetFail("查询失败")
	}
	return result.Result{
		Code: "0",
		Msg:  "查询成功",
		Data: json.RawMessage(data),
	}
}

func (service OrderService) GetOriROrders() []order.Order {
	res, err := orderMapper.GetOrders()
	if err != nil {
		return []order.Order{}
	}
	return res
}

func (service OrderService) GetPickedOrders(account string) *result.Result {
	m, err := orderMapper.RGetOrders()
	if err != nil {
		return result.GetFail("获取失败")
	}
	res := make([]string, len(m))
	for k, v := range m {
		if v == account {
			res = append(res, k)
		}
	}
	return &result.Result{
		Code: "0",
		Msg:  "获取成功",
		Data: res,
	}
}

func (service OrderService) ComOrder(uuid string) *result.Result {
	// MySql持久化
	if delROrders(uuid) {
		return result.GetFail("送达失败")
	}
	data := order.BroadcastData{
		OrderId: uuid,
		Flag:    "delete",
	}
	broadcast(uuid, data)
	return result.GetSuccess("成功送达")
}

func (service OrderService) AccOrder(uuid string, account string) *result.Result {
	flag, _ := orderMapper.IsPicked(uuid)
	if flag {
		return result.GetFail("该单已被接取")
	}
	err := orderMapper.SetOrderFlag(uuid, account)
	if err != nil {
		return result.GetFail("接单失败")
	}
	_ = riderSer.SetFlag(account, "false")
	data := order.BroadcastData{
		OrderId: uuid,
		Flag:    "accept",
	}
	broadcast(uuid, data)
	return result.GetSuccess("接单成功")
}

func (service OrderService) CanOrder(uuid string) *result.Result {
	flag, _ := orderMapper.IsPicked(uuid)
	if !flag {
		return result.GetFail("该单尚未被接取")
	}
	err := orderMapper.SetOrderFlag(uuid, "false")
	if err != nil {
		return result.GetFail("订单取消失败")
	}
	data := order.BroadcastData{
		OrderId: uuid,
		Flag:    "cancel",
	}
	broadcast(uuid, data)
	return result.GetSuccess("订单取消成功")
}

func (service OrderService) GetNums() int64 {
	return orderMapper.GetNum()
}

func (service OrderService) DelROrders(uuid string) *result.Result {
	if delROrders(uuid) {
		return result.GetFail("删除失败")
	}
	data := order.BroadcastData{
		OrderId: uuid,
		Flag:    "delete",
	}
	broadcast(uuid, data)
	return result.GetSuccess("删除成功")
}
