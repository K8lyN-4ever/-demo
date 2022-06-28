package mapper

import (
	"github.com/go-redis/redis"
	"rider/src/com/td/software/rider/common/resources"
	"rider/src/com/td/software/rider/common/util/order"
	"sync"
)

type Rider struct {
	id       int64  `gorm:"column:id"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	Type     int    `gorm:"column:type"`
}

func (Rider) TableName() string {
	return "orders.order"
}

type RiderMapper struct {
}

var riderMapper *RiderMapper
var riderOnce sync.Once

func NewRiderMapperInstance() *RiderMapper {
	riderOnce.Do(
		func() {
			riderMapper = &RiderMapper{}
		})
	return riderMapper
}

const (
	riderLocationKey = "rider_location"
	riderFlagKey     = "rider_flag"
	shopLocationKey  = "shop_location"
)

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

func delLocation(key string, name string) error {
	_, err := resources.RedisDb.ZRem(key, name).Result()
	if err != nil {
		return err
	}
	return nil
}

func getLocation(key string, name string) (order.Place, error) {
	res, err := resources.RedisDb.GeoPos(key, name).Result()
	if err != nil {
		return order.Place{}, err
	}
	if res[0] != nil {
		return order.Place{
			Longitude: res[0].Longitude,
			Latitude:  res[0].Latitude,
		}, nil
	}
	return order.Place{}, nil
}

func getTempByDist(key string, longitude float64, latitude float64, dist float64) ([]redis.GeoLocation, error) {
	res, err := resources.RedisDb.GeoRadius(key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:    dist,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
		Sort:      "ASC", //升序
	}).Result()
	if err != nil {
		return []redis.GeoLocation{}, err
	}
	return res, nil
}

func (mapper RiderMapper) SetLocation(name string, place order.Place) error {
	return addLocation(riderLocationKey, name, place)
}

func (mapper RiderMapper) SetFLag(name string, flag string) error {
	_, err := resources.RedisDb.HSet(riderFlagKey, name, flag).Result()
	if err != nil {
		return err
	}
	return nil
}

func (mapper RiderMapper) DelFlag(name string) error {
	_, err := resources.RedisDb.HDel(riderFlagKey, name).Result()
	if err != nil {
		return err
	}
	return nil
}

func (mapper RiderMapper) IsGrab(name string) bool {
	res, err := resources.RedisDb.HGet(riderFlagKey, name).Result()
	if err != nil {
		return false
	}
	return res == "true"
}

func (mapper RiderMapper) DelLocation(name string) error {
	return delLocation(riderLocationKey, name)
}

func (mapper RiderMapper) GetOrdersByDist(place order.Place) ([]redis.GeoLocation, error) {
	return getTempByDist(shopLocationKey, place.Longitude, place.Latitude, 3)
}

func (mapper RiderMapper) GetRidersByDist(place order.Place) ([]redis.GeoLocation, error) {
	return getTempByDist(riderLocationKey, place.Longitude, place.Latitude, 3)
}

func (mapper RiderMapper) GetRiderLocation(name string) (order.Place, error) {
	return getLocation(riderLocationKey, name)
}

func (mapper RiderMapper) GetOrderLocation(name string) (order.Place, error) {
	return getLocation(shopLocationKey, name)
}
