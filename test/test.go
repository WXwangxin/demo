package main

import (
	"fmt"
)

type Instance struct {
	Name string
	Data *int
}

func (i Instance) Store(num int) {
	*(i.Data) = num
}

func (i Instance) StoreName(name string) {
	i.Name = name
}

func (i Instance) Show() int {

	return *(i.Data)
}

func (i Instance) ShowName() string {

	return i.Name
}

/*
其实质依然是传参，在调用过程会把该对象作为第一个入参传入到函数里面，所以结构体，是值传递，copy一份对象，
对于结构体内的指针类型，由于copy后仍是保存的变量的内存地址，所以进行*操作可以影响外围对象
*/
func main() {
	data := 5

	i := &Instance{
		Name: "hello",
		Data: &data,
	}

	ref := i
	i.StoreName("iname")
	ref.Store(7)
	//ref.StoreName("store")

	fmt.Println(i.Show(), ref.Show())
	fmt.Println(i.ShowName(), ref.ShowName())
	// 打印出：7 7

	fmt.Println(fmt.Sprintf("%p %p", &i, &ref))
	// 打印出：0xc0000a6018 0xc0000a6030
}
