package main

import "fmt"

/*
Select permite que o programa tome decisão entre canais
A ordem não importa
tem default case igual o switch
Em momentos de competição o select escolhe o primeiro caso que não vai bloquear seja receber ou enviar
*/

func main() {

	//select são poderosos e permite que façamos design do fluxo de dados do programa.
	// neste caso os channels são tomados como "data streams"
	//select nos da o poder de gerenciar multiplas fontes de dados (streams)

	ch := make(chan int, 2)
	outA := make(chan int, 1)
	outB := make(chan int, 1)
	ch <- 1
	ch <- 1
	go Fanout(ch, outA, outB)

	if len(outA) > 0 {
		fmt.Println("a", <-outA)
	} else {
		fmt.Println("b", <-outB)
	}

	if len(outA) > 0 {
		fmt.Println("a", <-outA)
	} else {
		fmt.Println("b", <-outB)
	}

}

//Fan-out
//permite distribuirmos uma fonte de recebimento para diferentes leitores

func Fanout(in <-chan int, outA, outB chan int) {

	for d := range in { //recebe dados até o canal ser fechado

		select { //envia para o primeiro canal não bloqueado
		//situações em que os canais são "backup" um do outro;
		//tipo um load balancer; com multiplos workers na ponta fazendo o mesmo processo
		case outA <- d:
		case outB <- d:
		}

	}

}
