package mapper

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"rider/src/com/td/software/rider/common/resources"
	"rider/src/com/td/software/rider/common/util/order"
	"sync"
)

type Order struct {
	id       int64  `gorm:"column:id"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	Type     int    `gorm:"column:type"`
}

func (Order) TableName() string {
	return "orders.order"
}

type OrderMapper struct {
}

var orderMapper *OrderMapper
var orderOnce sync.Once

func NewOrderMapperInstance() *OrderMapper {
	orderOnce.Do(
		func() {
			orderMapper = &OrderMapper{}
		})
	return orderMapper
}

const (
	orderKey        = "orders"
	shopLocationKey = "shop_location"
	destLocationKey = "dest_location"
	orderFlagKey    = "orders_flag"
)

func (mapper OrderMapper) CreateOrder(order order.Order) error {
	return mapper.SetOrder(order)
}

func (mapper OrderMapper) SetOrder(order order.Order) error {
	value, err := json.Marshal(order)
	if err != nil {
		return err
	}
	resources.RedisDb.HSet(orderKey, order.Uuid, string(value))
	return nil
}

func (mapper OrderMapper) GetOrder(uuid string) (order.Order, error) {
	res, err := resources.RedisDb.HGet(orderKey, uuid).Result()
	if err != nil {
		return order.Order{}, err
	}
	result := order.Order{}
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return order.Order{}, err
	}
	return result, nil
}

func (mapper OrderMapper) GetOrders() ([]order.Order, error) {
	res, err := resources.RedisDb.HGetAll(orderKey).Result()
	if err != nil {
		return []order.Order{}, err
	}
	result := make([]order.Order, len(res))
	index := 0
	for _, v := range res {
		temp := order.Order{}
		err = json.Unmarshal([]byte(v), &temp)
		if err != nil {
			return []order.Order{}, err
		}
		result[index] = temp
		index += 1
	}
	return result, nil
}

func (mapper OrderMapper) DelOrder(uuid string) error {
	_, err := resources.RedisDb.HDel(orderKey, uuid).Result()
	if err != nil {
		return err
	}
	return nil
}

func addLocation(key string, name string, place order.Place) error {
	_, err := resources.RedisDb.GeoAdd(key, &redis.GeoLocation{
		Name:      name,
		Longitude: place.Longitude,
		Latitude:  place.Latitude,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}

func setOrderFlag(name string, flag string) error {
	_, err := resources.RedisDb.HSet(orderFlagKey, name, flag).Result()
	if err != nil {
		return err
	}
	return nil
}

func delLocation(key string, name string) error {
	_, err := resources.RedisDb.ZRem(key, name).Result()
	if err != nil {
		return err
	}
	return nil
}

func delOrderFlag(name string) error {
	_, err := resources.RedisDb.HDel(orderFlagKey, name).Result()
	if err != nil {
		return err
	}
	return nil
}

func (mapper OrderMapper) AddShopLocation(name string, place order.Place) error {
	return addLocation(shopLocationKey, name, place)
}

func (mapper OrderMapper) AddDestLocation(name string, place order.Place) error {
	return addLocation(destLocationKey, name, place)
}

func (mapper OrderMapper) DelShopLocation(name string) error {
	return delLocation(shopLocationKey, name)
}

func (mapper OrderMapper) DelDestLocation(name string) error {
	return delLocation(destLocationKey, name)
}

func (mapper OrderMapper) AddOrderFlag(name string) error {
	return setOrderFlag(name, "false")
}

func (mapper OrderMapper) SetOrderFlag(name string, flag string) error {
	return setOrderFlag(name, flag)
}

func (mapper OrderMapper) DelOrderFlag(name string) error {
	return delOrderFlag(name)
}

func (mapper OrderMapper) IsPicked(name string) (bool, error) {
	flag, err := resources.RedisDb.HGet(orderFlagKey, name).Result()
	if err != nil {
		return true, err
	}
	if flag == "false" {
		return false, nil
	}
	return true, nil
}

func (mapper OrderMapper) RGetOrders() (map[string]string, error) {
	m, err := resources.RedisDb.HGetAll(orderFlagKey).Result()
	if err != nil {
		return nil, err
	}
	return m, err
}

func (mapper OrderMapper) GetNum() int64 {
	result, err := resources.RedisDb.HLen(orderKey).Result()
	if err != nil {
		return 0
	}
	return result
}
