package main

import "fmt"

type Usb interface {
	// 声明了两个未实现的方法
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作。。。 ")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。 ")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。 ")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。 ")
}

type Computer struct {
}

// 接收一个接口类型变量，只要实现了Usb接口（实现接口声明的所有方法）
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 创建实现接口的结构体，传入接口结构体作为参数
	computer.Working(phone)
	computer.Working(camera)
}
