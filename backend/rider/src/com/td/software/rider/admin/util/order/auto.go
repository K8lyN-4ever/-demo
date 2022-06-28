package order

import "fmt"

func Start() {
	fmt.Println("自动分配协程开始")
	go dispatch()
	fmt.Println("自动生成协程开始")
	go generate()
}
