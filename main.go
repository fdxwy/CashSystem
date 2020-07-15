package main

import (
	"fmt"
	"CashSystem/strategy"

)


func main() {
	c := strategy.Context{A: 150, B: 7}
	// 用户自己决定使用什么策略
	c.SetContext(new(strategy.FullReduction))
	fmt.Println("总价为：",c.Result())
}
