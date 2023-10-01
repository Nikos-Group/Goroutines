package main

import (
	"fmt"
	"time"
)

func main() {
	var count int

	for i := 0; i < 1e5; i++ {
		/*
			Мы запускаем 100 тысяч горутин и ожидаем, что значение переменной
			count будет равно 100_000 к моменту ее вывода на экран
		*/
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println(count)
}
