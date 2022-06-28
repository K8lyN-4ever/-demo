package rider_service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	result "rider/src/com/td/software/rider/common/pojo"
	"rider/src/com/td/software/rider/common/util/order"
	"rider/src/com/td/software/rider/rider/mapper"
	"sync"
)

type RiderService struct {
}

var riderService *RiderService
var riderOnce sync.Once

func NewRiderServiceInstance() *RiderService {
	riderOnce.Do(
		func() {
			riderService = &RiderService{}
		})
	return riderService
}

var riderMapper = mapper.NewRiderMapperInstance()

func (service RiderService) SetLocation(name string, location order.Place) *result.Result {
	if err := riderMapper.SetLocation(name, location); err != nil {
		return result.GetFail("rider set location 数据库错误")
	}
	return result.GetSuccess("上传成功")
}

func (service RiderService) SetFlag(account string, flag string) *result.Result {
	if err := riderMapper.SetFLag(account, flag); err != nil {
		return result.GetFail("开启自动接单失败")
	}
	return result.GetSuccess("开启自动接单")
}

func (service RiderService) DelFlag(account string) *result.Result {
	if err := riderMapper.DelFlag(account); err != nil {
		return result.GetFail("删除骑手状态失败")
	}
	return result.GetSuccess("删除骑手状态")
}

func (service RiderService) IsGrab(account string) bool {
	return riderMapper.IsGrab(account)
}

func (service RiderService) DelLocation(name string) *result.Result {
	if err := riderMapper.DelLocation(name); err != nil {
		return result.GetFail("rider del location 数据库错误")
	}
	return result.GetSuccess("清除成功")
}

func (service RiderService) GetOriOrdersByDist(place order.Place) ([]redis.GeoLocation, error) {
	res, err := riderMapper.GetOrdersByDist(place)
	if err != nil {
		return []redis.GeoLocation{}, err
	}
	return res, nil
}

func (service RiderService) GetOriRidersByDist(place order.Place) ([]redis.GeoLocation, error) {
	res, err := riderMapper.GetRidersByDist(place)
	if err != nil {
		return []redis.GeoLocation{}, err
	}
	return res, nil
}

func (service RiderService) GetOrdersByDist(place order.Place) *result.Result {
	res, err := riderMapper.GetOrdersByDist(place)
	if err != nil {
		return result.GetFail("获取失败")
	}
	resStr, err := json.Marshal(res)
	return &result.Result{
		Code: "0",
		Msg:  "获取成功",
		Data: resStr,
	}
}

func (service RiderService) GetRiderLocation(name string) order.Place {
	res, err := riderMapper.GetRiderLocation(name)
	if err != nil {
		return order.Place{}
	}
	return res
}

func (service RiderService) GetOrderLocation(name string) order.Place {
	res, err := riderMapper.GetOrderLocation(name)
	if err != nil {
		return order.Place{}
	}
	return res
}
