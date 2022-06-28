package order

import "strconv"

type Place struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (p Place) String() string {
	longitudeStr := strconv.FormatFloat(p.Longitude, 'f', 18, 64)
	latitudeStr := strconv.FormatFloat(p.Latitude, 'f', 18, 64)
	return "Longitude" + longitudeStr + " " + "Latitude" + latitudeStr
}

type Order struct {
	id   int    `json:"id"`
	Uuid string `json:"uuid"`
	Src  Place  `json:"src"`
	Tar  Place  `json:"tar"`
}

type Rider struct {
	id      int    `json:"id" gorm:"id"`
	Account string `json:"account"`
	Pos     Place  `json:"pos"`
}

type BroadcastData struct {
	OrderId string `json:"order_id"`
	Flag    string `json:"flag"`
}
