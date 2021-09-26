package cncamp

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/**
  @params ch 双向通道，接收int型；
  @params wg 等待协程完成，注意这里传的是指针值（还是值传递）
  @params lock 互斥锁(sync.Mutex)
    互斥即不可同时运行。即使用了互斥锁的两个代码片段互相排斥，只有其中一个代码片段执行完成后，另一个才能执行。
*/
func producerA(pName string, ch chan int, wg *sync.WaitGroup, lock sync.Mutex) {
	go func(L sync.Mutex) {
		i := 1
		for {
			//L.Lock()
			ch <- i
			fmt.Println("p_in: ", pName+"_"+strconv.Itoa(i))
			// 等待10秒钟，如果想看明显的读阻塞，这里的沉睡时间就小很多倍，反之下面消费者里面就长很多倍；
			time.Sleep(time.Microsecond * 10)
			// 这里是控制到10次循环时就关闭通道
			if i == 10 {
				//close(ch)
				fmt.Println("finish")
				break
			}
			i++
			//L.Unlock()
		}
		fmt.Println("producer done")
		// 这句代表当前这次整体操作ok，不用再等待他了。
		wg.Done()
	}(lock)
}

/**
	@params: ch 双向通道 同上；
    @params: wg 同上
*/
func consumerA(ch chan int, wg *sync.WaitGroup) {
	go func() {
	OUT: // 指定break 跳出到for的外层
		for {
			select {
			case value := <-ch:
				fmt.Println("p_out", value)
				time.Sleep(time.Second * 1)
			case <-time.After(time.Microsecond * 5):
				fmt.Println("timeout")
				break OUT
			}
		}
		fmt.Println("consumer done")
		wg.Done()
	}()
}

// 多个生产者1个消费者
func TestMultiWriteOneRead(t *testing.T) {
	var wg sync.WaitGroup
	lock := sync.Mutex{}
	ch := make(chan int, 10)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		producerA("name_p"+strconv.Itoa(i), ch, &wg, lock)
	}

	wg.Add(1)
	consumerA(ch, &wg)

	wg.Wait()
}

// 多个生产者多个消费者
func TestMultiWriteMultiRead(t *testing.T) {
	var wg sync.WaitGroup
	lock := sync.Mutex{}

	for i := 0; i < 2; i++ {
		ch := make(chan int, 10)
		wg.Add(1)
		producerA("name_p"+strconv.Itoa(i), ch, &wg, lock)
		wg.Add(1)
		consumerA(ch, &wg)
	}

	wg.Wait()
}

// 注意：多个生产者共用一个channel时，不能在生产里面使用close(ch)关闭channel，否则会产生channel已经关闭的提示
// 因为 如果其中一个goroutine执行完毕，且执行close(ch)后，另一个channel是无法进行channel关闭的。
// 此时最好是采用select超时机制
