package websocket

import (
	"reflect"
	"strings"
)

type ControllerMapsType map[string]reflect.Value

func (manager *Manager) ServerCodeToFunc(data ReadData) {
	funcName := case2Camel(data.Actioncode)
	vft := manager.serverReturnFunc()
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(data)
	if vft[funcName].IsValid() {
		vft[funcName].Call(params)
	}
}

func case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func (manager *Manager) serverReturnFunc() ControllerMapsType {
	var m ServerMethod
	vf := reflect.ValueOf(&m)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	crMap := make(ControllerMapsType, 0)

	//遍历所有的方法，并将其存入映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

func (w *receiver) ClientCodeToFunc(data baseMsg) {
	funcName := case2Camel(data.Actioncode)
	vft := w.serverReturnFunc()

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(data)
	if vft[funcName].IsValid() {
		vft[funcName].Call(params)
	}
}

func (w *receiver) serverReturnFunc() ControllerMapsType {
	var m ClientMethod
	vf := reflect.ValueOf(&m)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	crMap := make(ControllerMapsType, 0)

	//遍历所有的方法，并将其存入映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}
