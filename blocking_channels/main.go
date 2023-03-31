package main

import (
	"fmt"
	"time"
)

func main() {

	//blocking channels

	buffer := make(chan int, 1)

	/*
		//lendo o canal vazio = no data is blocking!
		fmt.Println(<-buffer)
	*/

	//guarda envio 1 :
	buffer <- 1
	//proximo envio está blocado.
	buffer <- 2

	//agora sincronizamos sender e receiver.
	go func() {

		for a := range buffer {
			fmt.Println(a)
		}

	}()
	buffer <- 2

	time.Sleep(time.Second * 5)
}
