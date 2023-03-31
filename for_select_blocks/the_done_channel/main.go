package main

import (
	"fmt"
	"time"
)

// este padrão serve para ajudar a previnir que go routines
//executem para além do tempo necessário
//a rotina parent fecha a child

func main() {

	done := make(chan bool)

	go processData(done)

	time.Sleep(time.Second * 2)

	close(done)
	time.Sleep(time.Second * 2)

}

func processData(done <-chan bool) {

	for {
		select {
		case <-done: //caso o canal seja fechado indicara que a rotina precisa ser fechada
			fmt.Println(" ---- Rotina será fechada ---- ")
			return
		default:
			fmt.Println("PROCESSING")
		}
	}

}
