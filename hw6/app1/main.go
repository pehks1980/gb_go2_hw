package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

/*
Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
Выполните трассировку программы
*/

/*
программа использует множество SET, поточек i генерит СЧ проверяет есть ли такое число i == в множестве
если да то ф. меняет его с СЧ, иначе берется это число и к нему прибавляется CЧ и заводится новое число в это множество
*/

// наша мапа множество
type Set struct {
	sync.Mutex
	mm map[int64]int
}

// конструктор множества
func NewSet() *Set {
	return &Set{mm: make(map[int64]int)}

}

// добавление числа i в множество флаг 1
func (s *Set) Add(i int) {
	s.Lock()
	s.mm[int64(i)] = 1
	s.Unlock()
}

// добавление числа i в множество задано без семафора
func (s *Set) AddNoLock(i int) {
	//s.Lock()
	s.mm[int64(i)] = 1
	//s.Unlock()
}

// удаление из множества (помечает как 0)
func (s *Set) Del(i int) {
	s.Lock()
	s.mm[int64(i)] = 0
	s.Unlock()
}

// метод проверки есть число i в множестве или нет
func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	val, ok := s.mm[int64(i)]
	if val == 1 && ok {
		return true
	}
	return false
}

// функция генерит число в диапазоне
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// функция обработки поточка
func checkNMultiply(i int, set *Set) error {
	//get random
	randInt := randInt(5, 100)
	//check if we have i in set
	if set.Has(i) {
		set.Del(i)
		if RaceOn {
			set.AddNoLock(randInt)
		} else {
			set.Add(randInt)
		}

		fmt.Printf("number %d in Set! replaced with %d\n", i, randInt)
	} else {
		if RaceOn {
			set.AddNoLock(randInt + i)
		} else {
			set.Add(randInt + i)
		}

		fmt.Printf("number %d is not Set! new number added %d\n", i, randInt+i)
	}
	return nil
}

var (
	// флаг модифицирует код который исполняет добавление в множество без семафора
	// методом AddNoLock
	// запуск и листинг go run -race main.go 2> race.out
	RaceOn bool = false
	// флаг запускает трассировку
	// cделать и поглядеть трассировку:
	// GOMAXPROCS=1 go run main.go > trace.out
	// go tool trace trace.out
	TraceOn bool = true
	ShedOn bool = false
)
// N Число горутин
const N int = 100

func main() {

	// trace code
	if TraceOn {
		trace.Start(os.Stderr)
		defer trace.Stop()
	}

	var set = NewSet()
	var wg = sync.WaitGroup{}
	rand.Seed(time.Now().UTC().UnixNano())
	// N Число горутин

	for i := 1; i <= N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := checkNMultiply(i, set)
			if err != nil {
				log.Printf("some error occured %v", err)
			}
		}(i)
	}
	// вызов планировщика на перепланировку
	go func() {
		for i := 0; i < 5; i += 1 {
			time.Sleep(1 * time.Millisecond)
			if ShedOn {
				runtime.Gosched()
				fmt.Println("Gosсhed")
			}
		}
	}()

	//ждем пока все обработается
	wg.Wait()

	fmt.Println("завершено")

	// распечатываем ресурс
	for k, v := range set.mm {
		if true {
			fmt.Printf("%d -> %d\n", k, v)
		}

	}

}
