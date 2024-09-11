package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	NoPool()
	CommonPool()
	Pool()
	PoolPreAlloc()
	PoolWithFunc()
	MultiPool()
	MultiPoolWithFunc()
}

var (
	runTimes = 1000
	sum      int32
	wg       sync.WaitGroup
)

func Add(i any) {
	num := i.(int32)
	atomic.AddInt32(&sum, num)
	fmt.Println(" run with ", num)
}

func test() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func NoPool() {
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		go func() {
			test()
			wg.Done()
		}()
	}
	wg.Wait()
}
func CommonPool() {
	defer ants.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(func() {
			test()
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println("running goroutines:", ants.Running())
}

func Pool() {
	p, _ := ants.NewPool(runTimes)
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Submit(func() {
			test()
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println("running goroutines:", ants.Running())
}

func PoolPreAlloc() {
	//预先分配 goroutine 队列内存
	p, _ := ants.NewPool(runTimes, ants.WithPreAlloc(true))
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Submit(func() {
			test()
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println("running goroutines:", ants.Running())
}

func PoolWithFunc() {
	p, _ := ants.NewPoolWithFunc(100, func(i any) {
		Add(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Println("running goroutines: ", p.Running())
	fmt.Println("finish all tasks, result is ", sum)
}

func MultiPool() {
	// Use the MultiPool and set the capacity of the 10 goroutine pools to unlimited.
	// If you use -1 as the pool size parameter, the size will be unlimited.
	// There are two load-balancing algorithms for pools: ants.RoundRobin and ants.LeastTasks.
	mp, _ := ants.NewMultiPool(10, -1, ants.RoundRobin)
	defer mp.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		mp.Submit(func() {
			test()
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println("running goroutines: ", mp.Running())
}

func MultiPoolWithFunc() {
	// Use the MultiPoolFunc and set the capacity of 10 goroutine pools to (runTimes/10).
	mpf, _ := ants.NewMultiPoolWithFunc(10, 10, func(i any) {
		Add(i)
		wg.Done()
	}, ants.LeastTasks)
	//5秒后释放资源
	defer mpf.ReleaseTimeout(5 * time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mpf.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Println("running goroutines: ", mpf.Running())
}
