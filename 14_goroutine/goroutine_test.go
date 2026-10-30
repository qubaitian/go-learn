package demo

import (
	"sync"
	"sync/atomic"
	"testing"
)

// go 关键字启动一个 goroutine（轻量级线程）；用 WaitGroup 等它们全部结束
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	results := make([]int, 5)

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = i * i // 每个 goroutine 写不同的下标，所以不冲突
		}()
	}
	wg.Wait() // 坑：不 Wait 的话，测试函数可能在 goroutine 跑完之前就返回了

	t.Log(results)
	if results[4] != 16 {
		t.Error("结果不对", results)
	}
	// 注意：Go 1.22 之前，循环变量 i 是所有迭代共用的，闭包里会读到错的值；1.22+ 每次迭代都是新的 i
}

// wg.Go（Go 1.25+）：把 Add / go / Done 合成一步，更简洁
func TestWaitGroupGo(t *testing.T) {
	var wg sync.WaitGroup
	var n atomic.Int64

	for range 100 {
		wg.Go(func() { n.Add(1) }) // atomic 原子操作，并发加也是安全的
	}
	wg.Wait()

	if n.Load() != 100 {
		t.Errorf("期望 100，实际 %d", n.Load())
	}
}

// 坑：多个 goroutine 同时读写同一个变量就是数据竞争，结果不可预测
// 用 Mutex 保护共享数据。用 go test -race 可以检测出数据竞争
func TestMutex(t *testing.T) {
	var (
		mu    sync.Mutex
		count int
		wg    sync.WaitGroup
	)

	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			count++
		})
	}
	wg.Wait()

	if count != 1000 {
		t.Errorf("期望 1000，实际 %d", count)
	}
}

// sync.Once：不管调用多少次、多少个 goroutine 同时调用，都只执行一次（常用于懒加载）
func TestOnce(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup
	calls := 0

	for range 10 {
		wg.Go(func() {
			once.Do(func() { calls++ })
		})
	}
	wg.Wait()

	if calls != 1 {
		t.Errorf("期望只执行 1 次，实际 %d 次", calls)
	}
}
