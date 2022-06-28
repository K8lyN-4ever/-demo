package websocket

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type conn struct {
	c   *websocket.Conn
	mtx *sync.Mutex
}

var (
	addr   string
	client = conn{mtx: new(sync.Mutex), c: nil}

	httpReqMsgReceiver = receiver{
		cmdMtx:    new(sync.Mutex),
		curMaxCmd: 0,

		receiverMtx:       new(sync.RWMutex),
		maxReceiver:       100,
		receiver:          make(map[int]chan []byte),
		receiveMsgTimeout: 6 * time.Second,
	}

	clientArr = make(map[string]*websocket.Conn)
)

type receiver struct {
	curMaxCmd int
	cmdMtx    *sync.Mutex

	receiver    map[int]chan []byte
	receiverMtx *sync.RWMutex
	maxReceiver int

	receiveMsgTimeout time.Duration
}

func (w *receiver) pushMap(k int, c chan []byte) int {
	w.receiverMtx.Lock()
	if _, ok := w.receiver[k]; ok {
		return 51003
	}
	w.receiver[k] = c
	w.receiverMtx.Unlock()

	return 200
}

func (w *receiver) deleteMap(k int) {
	w.receiverMtx.Lock()
	if _, ok := w.receiver[k]; ok {
		if len(w.receiver[k]) > 0 {
			<-w.receiver[k]
		}
		close(w.receiver[k])
		delete(w.receiver, k)
	}
	w.receiverMtx.Unlock()
}

func (w *receiver) getCmd() (v int) {
	w.cmdMtx.Lock()
	w.curMaxCmd += 1
	v = w.curMaxCmd
	w.cmdMtx.Unlock()
	return
}

func (w *receiver) recoverCmd(v int) {
	if len(w.receiver) == 0 {
		w.cmdMtx.Lock()
		w.curMaxCmd = 0
		w.cmdMtx.Unlock()
		return
	}

	if v == w.curMaxCmd {
		w.cmdMtx.Lock()
		w.curMaxCmd -= 1
		w.cmdMtx.Unlock()
	}
}

type baseMsg struct {
	Company       string      `json:"company"`
	Actioncode    string      `json:"actioncode"`
	Data          interface{} `json:"data"`
	ModID         string      `json:"modID"`
	Token         string      `json:"token"`
	Result        int         `json:"result"`
	ResultMessage string      `json:"result_message"`
	CmdSequence   int         `json:"CmdSequence"`
}

//发送给固定客户端
func (w *receiver) receiveMsg(cmd int, msg []byte) {
	w.receiverMtx.RLock()
	defer w.receiverMtx.RUnlock()
	fmt.Println(9)
	if _, ok := w.receiver[cmd]; ok {
		go func() {
			select {
			case w.receiver[cmd] <- msg:
			case <-time.After(1 * time.Second):
			}
		}()
	}
}
