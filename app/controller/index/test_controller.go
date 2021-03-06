package index

import (
	"fmt"
	"github.com/wycto/weigo"
)

type TestController struct {
	weigo.Controller
}

func (receiver *TestController) Index() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func (receiver *TestController) Route() {
	fmt.Println("param", receiver.Context.Param)
	fmt.Println("Get", receiver.Context.Get)
	fmt.Println("Post", receiver.Context.Post)
	fmt.Println("---------")
}
