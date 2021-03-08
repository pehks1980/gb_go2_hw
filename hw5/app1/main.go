package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

/*
Напишите программу, которая запускает  потоков и дожидается завершения их всех
*/

var (
	Net string = "192.168.1."
)

//Функция пингует все хосты сетки, запуская в поточке команду ping
func pingNet(n int) error {

	app := "ping"
	arg0 := Net + strconv.Itoa(n)
	arg1 := "-c"
	arg2 := "4"
	// пингует 4 пингами
	cmd := exec.Command(app, arg1, arg2, arg0)
	fmt.Println("pinging ", cmd)

	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(stdout))
	return nil
}

func main() {

	var wg = sync.WaitGroup{}
	// Число хостов в сетке
	N := 254

	for i := 1; i <= N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := pingNet(i)
			if err != nil {
				log.Printf("host %s%d is not responding...", Net, i)
			}
			//wg.Done()
		}(i)
	}

	//ждем пока все перемножится
	wg.Wait()

	fmt.Println("Пингование завершено")

}
