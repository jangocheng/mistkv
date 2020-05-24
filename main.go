package main

import (
	"fmt"
	"mistkv/buffer"
	"mistkv/kv"
)

func main() {
	pol := kv.CreatePool()
	rds := pol.Get()
	defer rds.Close()

	chanBuffer := buffer.NewChanBuffer()
	// kv.SetIncreasLast(rds, "mistLast", 500)  // 设置最大值
	// mistLast, err := kv.GetIncreasLast(rds, "mistLast") // 取出最大值
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mistLast)
	// 查看余量
	// sur, err := kv.SurplusMistValue(rds, "mistList")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("sur: ", sur)

	// 写入值
	// for i := 20; i < 50; i++ {
	// 	kv.LpushMistValue(rds, "mistList", i)
	// }
	fmt.Println("Before chanBuffer len is: ", len(chanBuffer.Channel))
	// 取出值
	for i := 0; i < 10; i++ {
		val, err := kv.RpopMistValue(rds, "mistList")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
			chanBuffer.Put(val)
		}
	}
	fmt.Println("After chanBuffer len is: ", len(chanBuffer.Channel))
	// 查看余量
	sur, err := kv.SurplusMistValue(rds, "mistList")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sur: ", sur)
	fmt.Println("welcome to mist server")
}
