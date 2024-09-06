package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	//创建observable
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	ch := observable.Observe()
	for item := range ch {
		fmt.Println("observable1:", item.V)
	}

	ch2 := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- rxgo.Of(i)
		}
		close(ch2)
	}()

	// FromChannel可以直接从一个已存在的<-chan rxgo.Item对象中创建 Observable
	observable2 := rxgo.FromChannel(ch2)
	for item := range observable2.Observe() {
		fmt.Println("observable2", item.V)
	}

	//Interval以传入的时间间隔生成一个无穷的数字序列，从 0 开始
	observable3 := rxgo.Interval(rxgo.WithDuration(5 * time.Second))
	for item := range observable3.Observe() {
		fmt.Println("observable3", item.V)
	}

	//Range可以生成一个范围内的数字
	observable4 := rxgo.Range(0, 3)
	for item := range observable4.Observe() {
		fmt.Println("observable4", item.V)
	}

	//Observable 对象上调用Repeat，可以实现每隔指定时间
	//重复一次该序列，一共重复指定次数
	observable5 := rxgo.Just(1, 2, 3)().Repeat(
		3, rxgo.WithDuration(1*time.Second),
	)
	for item := range observable5.Observe() {
		fmt.Println(item.V)
	}
}
