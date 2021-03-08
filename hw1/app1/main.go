package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/*
Напишите программу, в которой неявно будет срабатывать паника. Сделайте отложенную функцию,
которая будет обрабатывать эту панику и печатать предупреждение в консоль.
Критерий выполнения задания — программа не завершается аварийно. ё
Дополните программу собственной ошибкой, хранящей время её возникновения.

Напишите функцию, которая создаёт файл в файловой системе и использует отложенный вызов функций
для безопасного закрытия файла.
*/

//структура кастомной ошибки
type ErrorWithTime struct {
	text     string
	dateTime string
}

// функция создания для кастомной ошибки которая еще печатает дату
func New(text string) error {
	t := time.Now()
	return &ErrorWithTime{
		text:     text,
		dateTime: t.String(),
	}
}

// Задание своего метода Error для интерфейса ошибки error своя структура ErrorWithTime
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s time: %s\n", e.text, e.dateTime)
}

// функция деления 2 чисел float с обработкой исключения при делении на 0
func customDivide(numberA float64, numberB float64) float64 {
	// recovery for panic sit defer - starts BEFORE actual PANIC!!!! аналог except: python
	defer func() {
		err := New("divide by zero.")
		if v := recover(); v != nil {
			fmt.Println("Делитель не должен быть по идее = 0! ", v)
			fmt.Println(err)
		}
	}()

	if numberB == 0.0 {
		// вызываем панику т.к. в случае с float результат = +inf
		panic("ОШИБКА делитель = 0!")
	}

	return numberA / numberB
}

// Функция записывает в файл 3 операнда а, б и результат целочисл. деления
func dumpToFile(fileName string, nA, nB, res float64) error {
	// Пробуем создать файл
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	n, err := fmt.Fprintf(file, "Делитель=%.4f\n", nA)
	if err != nil {
		return err
	}
	n, err = fmt.Fprintf(file, "Делимое=%.4f\n", nB)
	if err != nil {
		return err
	}
	n, err = fmt.Fprintf(file, "Частное=%.4f\n", res)
	if err != nil {
		return err
	}

	fmt.Println(n, "bytes written")
	fmt.Println("done")

	return nil
}

func main() {

	aNumber := 3.5
	bNumber := 0.0

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	filename := path + "/hw1/app1/" + "out.txt"

	flag.Parse()
	filename1 := strings.TrimSpace(flag.Arg(0)) // Считываем имя файла и очищаем ввод от пробелов

	if filename1 != "" {
		// если прямо задать аргументом имя файла, оно будет использоваться
		filename = filename1
	}

	result := customDivide(aNumber, bNumber)

	// функция создает файл и использует отложенный вызов для безопасного закрытия файла
	err = dumpToFile(filename, aNumber, bNumber, result)
	if err != nil {
		log.Printf("Ошибка файла %v", err)
	}

	fmt.Println("result=", result)
}
