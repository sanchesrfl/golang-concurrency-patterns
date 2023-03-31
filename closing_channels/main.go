package main

//sempre que fechamos um channel produzimos uma mensagem;

//quando enviamos algo para um canal fechado, go routine panics!

import "fmt"

func main() {

	c := make(chan int)
	close(c)
	fmt.Print(<-c) //o receiver sabe sempre que o canal fechou.
	//mas o sender não. Portanto, é melhor fechar sempre do lado do sender.

	//simular sender panic
	c <- 1

}
