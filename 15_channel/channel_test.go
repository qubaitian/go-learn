package demo

import (
	"context"
	"slices"
	"testing"
	"time"
)

// 无缓冲 channel：发送方会一直阻塞，直到有人接收（相当于一次"握手"）
// Go 的理念：不要通过共享内存来通信，而要通过通信来共享内存
func TestUnbuffered(t *testing.T) {
	ch := make(chan string)

	go func() { ch <- "hello" }() // 发送
	msg := <-ch                   // 接收

	if msg != "hello" {
		t.Error("没收到消息")
	}
	// 坑：在同一个 goroutine 里对无缓冲 channel 先发后收会死锁 (fatal error: all goroutines are asleep)
}

// 有缓冲 channel：缓冲没满之前，发送不会阻塞
func TestBuffered(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // 缓冲是 2，所以这两次发送都不会阻塞

	if len(ch) != 2 || cap(ch) != 2 {
		t.Error("len/cap 不对")
	}
}

// close + range：发送方用完了就 close，接收方用 range 一直收到 channel 被关闭为止
// 参数里的 chan<- int 表示"只能发送"的 channel，<-chan int 表示"只能接收"
func produce(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // 坑：只有发送方才应该 close；向已关闭的 channel 发送会 panic
		for i := range n {
			ch <- i
		}
	}()
	return ch
}

func TestCloseRange(t *testing.T) {
	var got []int
	for v := range produce(3) {
		got = append(got, v)
	}

	if !slices.Equal(got, []int{0, 1, 2}) {
		t.Error("结果不对", got)
	}
}

// 从已关闭的 channel 接收：立刻返回零值，ok=false
func TestReceiveFromClosed(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 42
	close(ch)

	v1, ok1 := <-ch // 缓冲里还有数据，照常收到
	v2, ok2 := <-ch // 已经空了并且关闭了

	if v1 != 42 || !ok1 || v2 != 0 || ok2 {
		t.Error("结果不对", v1, ok1, v2, ok2)
	}
}

// select：同时等多个 channel，哪个先准备好就走哪个分支；常用来做超时控制
func TestSelectTimeout(t *testing.T) {
	slow := make(chan string, 1) // 带 1 个缓冲，超时后没人接收时 goroutine 也能发完退出，不会泄漏
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "done"
	}()

	select {
	case msg := <-slow:
		t.Error("不应该先收到", msg)
	case <-time.After(50 * time.Millisecond):
		t.Log("超时了")
	}
}

// select + default：不阻塞，有数据就收，没有就直接走 default
func TestSelectDefault(t *testing.T) {
	ch := make(chan int)

	select {
	case v := <-ch:
		t.Error("不应该收到", v)
	default:
		t.Log("没有数据，直接走 default")
	}
}

// context：用来通知一组 goroutine "该停下了"（超时 / 取消），实际项目里到处都会用到
func worker(ctx context.Context, out chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			close(out)
			return
		case out <- i:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	out := make(chan int)
	go worker(ctx, out)

	count := 0
	for range out {
		count++
	}

	t.Log("超时前收到了", count, "个数")
	if count == 0 {
		t.Error("应该收到一些数据")
	}
}
