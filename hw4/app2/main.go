package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
позднее, чем за одну секунду (установить таймаут).

*/
var (
	// управление - сигналом
	SIGTERM bool = true //false
	signal       = make(chan bool)
)

// функция остановки, завершает программу не позже 1 с.
func shutDown() error {
	//Создаёт новый контекст, который завершится максимум через секунду
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	// т.е через 1 секунду будет завершена по любому
	defer cancelFunc()

	doneCh := make(chan error)
	go func(ctx context.Context) {
		// запуск процедуры завершения
		err := shutDownNow(ctx)
		doneCh <- err
	}(ctx)

	var err error
	// селект блокирует функцию которая завершится либо по ctx.Done которая приходит по cancel
	// второй случай shutDownNow вернула какое то значение ошибки или nil ошибок
	// ctx.Err это ошибка контекста - сработал таймаут
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-doneCh:
	}
	return err
}

//Функция делающая действия (эмулирующая) аварийную остановку по Sigterm
func shutDownNow(ctx context.Context) error {
	time.Sleep(3 * time.Second)
	// ...
	return nil
}

func main() {
	//посылает сигнал SIGTERM в канал и закрывает его
	go func() {
		if SIGTERM {
			signal <- SIGTERM

		}
		close(signal)
	}()

	for {
		_, ok := <-signal
		if !ok {
			// если канал закрыли и не было сигнала SigTERM ничего не делаем.
			fmt.Println("Нормальное завершение/выполнение. NO SIGTERM!.")
			return
		}

		select {
		case <-signal:
			// в случае SIGTERM выполняем логику shotDown
			err := shutDown()
			if err != nil {
				log.Printf("SIGTERM - Остановка, сработал тайм аут: %v", err)
			} else {
				log.Printf("SIGTERM - Остановка, OK\n")
			}
			return
		}

	}

}
