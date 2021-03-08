package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
С помощью пула воркеров написать программу, которая запускает 1000 горутин,
каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться,
что при каждом запуске программы итоговое число равно 1000.

*/
var (
	counter      int64 = 1
	max_counter  int64 = 1000
	ctrl_channel       = make(chan struct{})
	// канал workers буфф
	workers = make(chan struct{}, max_counter)
)

func main() {
	// отлов значения каунтера
	go func() {
		// цикл for будет увеличиваться пока не пройдет 100 сек макс
		for i := 1; i < 1000; i++ {

			if counter > 1000 {
				//put control channel stop
				// по достижении каунтера 1000 кидаем указатель в канал управления
				ctrl_channel <- struct{}{}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//запускаем 1000 поточков
	for i := 1; i <= int(max_counter); i++ {
		//поточек обрабатывает
		go func(job int) {

			defer func() {
				//после отработки извлекаем 1 элемент из канала workers
				<-workers
			}()
			atomic.AddInt64(&counter, 1)
			fmt.Printf("go job id =%d has finished and increased counter by 1 to value %d\n", job, counter)
			// спит есть работает делает что то еще...
			time.Sleep(1 * time.Second)
		}(i)
	}

	// ждем сигнала окончания
	select {
	case <-ctrl_channel:
		fmt.Println("all done.")
	}

	fmt.Println("main thread finished.")

}
