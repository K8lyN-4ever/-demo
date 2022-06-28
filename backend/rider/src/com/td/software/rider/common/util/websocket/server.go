package websocket

import (
	"encoding/json"
	"fmt"
	"rider/src/com/td/software/rider/rider/service"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string
	Socket  *websocket.Conn
	Message chan []byte
}

type BroadCastMessageData struct {
	Message     []byte
	IsBroadCast bool
	ClientIDs   []string
}

type Manager struct {
	Group                map[string]*Client
	Lock                 sync.Mutex
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
	clientCount          uint //分组及客户端数量
}

var WebsocketManager = Manager{
	Group:            make(map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	clientCount:      0,
}

func (manager *Manager) Start() {
	fmt.Println("websocket 服务器启动")
	for {
		select {
		case client := <-manager.Register:
			//注册客户端
			manager.Lock.Lock()
			manager.Group[client.Id] = client
			manager.clientCount += 1
			fmt.Printf("客户端注册: 客户端id为%s\n", client.Id)
			manager.Lock.Unlock()
			//_ = client.Socket.WriteMessage(websocket.TextMessage, res)
		case client := <-manager.UnRegister:
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Id]; ok {
				close(client.Message)
				delete(manager.Group, client.Id)
				manager.clientCount -= 1
				res := rider_service.NewRiderServiceInstance().DelLocation(client.Id)
				_ = rider_service.NewRiderServiceInstance().DelFlag(client.Id)
				ids := make([]string, 1)
				ids[0] = client.Id
				manager.Success(res, false, ids)
				fmt.Printf("客户端注销: 客户端id为%s\n", client.Id)

			}
			manager.Lock.Unlock()
		case data := <-manager.BroadCastMessage:
			for _, conn := range manager.Group {
				if data.IsBroadCast {
					conn.Message <- data.Message
				} else {
					for _, id := range data.ClientIDs {
						if id == conn.Id {
							conn.Message <- data.Message
							break
						}
					}

				}

			}

		}
	}
}

type ReadData struct {
	Actioncode  string   `json:"actioncode"`
	Data        string   `json:"data"`
	Location    string   `json:"location"`
	IsBroadCast bool     `json:"is_broadcast"`
	ClientIDs   []string `json:"client_ids"`
}

func (c *Client) Read() {
	defer func() {
		if err := c.Socket.Close(); err != nil {
			fmt.Printf("client [%s] disconnect err: %s\n", c.Id, err)
		}
		WebsocketManager.UnRegisterClient(c)
		fmt.Printf("client [%s],客户端关闭：[%v]\n", c.Id, websocket.CloseMessage)
	}()

	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			fmt.Printf("client [%s],数据读取失败或通道关闭：[%s],客户端连接状态：[%v]\n", c.Id, err.Error(), websocket.CloseMessage)
			break
		}
		var data ReadData
		err = json.Unmarshal(message, &data)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("数据解析失败\n")
		}

		data.ClientIDs = append(data.ClientIDs, c.Id)
		WebsocketManager.ServerCodeToFunc(data)

	}
}

func (c *Client) Write() {
	defer func() {
		if err := c.Socket.Close(); err != nil {
			//log.WSLog(fmt.Sprintf("client [%s] disconnect err: %s", c.Id, err))
			return
		}
		WebsocketManager.UnRegisterClient(c)
		fmt.Printf("client [%s],客户端关闭：[%v]\n", c.Id, websocket.CloseMessage)
	}()
	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				fmt.Printf("client [%s],客户端连接状态：[%v]\n", c.Id, websocket.CloseMessage)
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				//log.WSLog(fmt.Sprintf("client [%s] write message err: %s", c.Id, err))
				return
			}
		}
	}
}

func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

type ResultData struct {
	Company       string      `json:"company"`
	DeviceName    string      `json:"device_name"`
	Result        int         `json:"result"`
	ResultMessage string      `json:"result_message"`
	Version       string      `json:"version"`
	DBVersion     string      `json:"db_version"`
	Language      string      `json:"language"`
	Data          interface{} `json:"data"`
	ActionCode    string      `json:"actioncode"`
}

func (manager *Manager) Success(data interface{}, isBroadCast bool, ClientIDs []string) {
	msg, err := json.Marshal(data)
	if err != nil {
		return
	}
	WebsocketManager.BroadCastMessage <- &BroadCastMessageData{Message: msg, IsBroadCast: isBroadCast, ClientIDs: ClientIDs}
}

func (manager *Manager) Error(data []interface{}, isBroadCast bool, ClientIDs []string) {
	msg, err := json.Marshal(data)
	if err != nil {
		return
	}
	WebsocketManager.BroadCastMessage <- &BroadCastMessageData{Message: msg, IsBroadCast: isBroadCast, ClientIDs: ClientIDs}
}
