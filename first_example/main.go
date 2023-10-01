package main

import (
	"fmt"
	"time"
)

func main() { // это такая же горутина, как и те которые мы вызываем
	/*
		как толкьо завершается "главная" горутина,
		то завершаются и все дочерние горутины
	*/
	go func() {
		fmt.Println("Hello from anonymous func")
		time.Sleep(time.Second * 2)
		fmt.Println("hello 2")
	}()

	go printhello() // через именованную функцию

	var p printer // определяем структуру
	go p.printhello()

	time.Sleep(time.Second)
}

func printhello() {
	fmt.Println("Hello from named func")
}

type printer struct {
}

func (printer) printhello() {
	fmt.Println("hello from struct method")
}
