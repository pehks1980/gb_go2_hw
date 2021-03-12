package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализуйте функцию для разблокировки мьютекса с помощью defer
 */

const count = 1000

// функция критич. секции
// есть разблокировка мьютекса с помощью defer + обработка паники
func criticalSection(mut *sync.Mutex, ctr *int64) {
	defer func() {
		fmt.Println("recovered", recover())
	}()
	mut.Lock()
	defer mut.Unlock()
	*ctr += 1
	panic("AAA!")
}

func main() {
	var (
		counter int64 = 0
		mutex sync.Mutex

		// Вспомогательная часть нашего кода
		ch = make(chan struct{}, count)
	)
	// запускаем 1000 горутин в каждая из которых увеличивает каунтер и заносит пустую структур в канал
	for i := 0; i < count; i += 1 {
		go func() {
			// функция критич. секции
			// там есть разблокировка мьютекса с помощью defer
			criticalSection(&mutex,&counter)

			// Фиксация факта запуска горутины в канале
			ch <- struct{}{}
		}()
	}
	/// после запуска ждем 2 сек и закрываем канал
	time.Sleep(2*time.Second)
	close(ch)
	//считаем сколько горутин выполнялось - сколько структур накидалось в канал
	i := 0
	for range ch {
		i += 1
	}
	// Выводим показание счетчика
	fmt.Println(counter)
	// Выводим показания канала
	fmt.Println(i)
	// в итоге должно быть одно число.
}