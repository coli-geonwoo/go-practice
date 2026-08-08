//고루틴 동기화 기초(5)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//Wait , notify , notifyAll : 기타 언어
	//Wait , Signal , Broadcast

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock() //락 취득
			c <- 777
			fmt.Println("Goroutine Wating : ", n)
			//Wait() 내부 동작: mutex.Unlock() -> 대기(blocking) -> 깨어나면 mutex.Lock()
			//즉 대기에 진입하는 순간 락을 자동으로 놓아주므로, 이 시점부터 다른 고루틴/메인이 락 취득 가능
			condition.Wait() //고루틴 대기(멈춤) -> 획득한 락을 해제하고 대기하고 있다가 깨어났을 때 다시 취득
			//Signal/Broadcast로 깨어난 뒤, Wait() 내부에서 락을 다시 잡은 상태로 여기 도달
			fmt.Println("Wating End : ", n)
			mutex.Unlock() //Wait()가 재취득해준 락을 해제 -> 다음 고루틴이 락을 잡고 진행
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("received : ", <-c)
	}

	//[Signal 방식] 대기자 중 "1개"만 깨움 -> 5개를 깨우려면 5번 반복 필요
	//for i := 0; i < 5; i++ {
	//	mutex.Lock()
	//	fmt.Println("Wake Goroutine(Signal) : ", i)
	//	condition.Signal() //한 개 씩 깨움(모든 고루틴 생성 후)
	//	mutex.Unlock()
	//}

	//[Broadcast 방식] 대기 중인 "모든" 고루틴을 한 번의 호출로 깨움
	mutex.Lock() //메인이 락 취득
	fmt.Println("Wake Goroutine(Broadcast)")
	//Broadcast는 대기 고루틴들을 "실행 가능" 상태로 표시만 할 뿐, 락을 넘겨주지 않음
	//깨어난 고루틴들은 Wait() 내부의 mutex.Lock()에서 다시 막혀 있는 상태
	condition.Broadcast()
	//메인이 락을 놓아야 비로소 깨어난 고루틴들이 하나씩 락을 잡으며 순차적으로 진행됨
	mutex.Unlock()

	time.Sleep(2 * time.Second)

}
