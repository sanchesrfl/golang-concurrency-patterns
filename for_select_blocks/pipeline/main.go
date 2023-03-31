package main

import "fmt"

func main() {

	//pipelines são formas de organizarmos o fluxo de dados em estágios sincronizados

	//input //start pipeline
	n := []int{3, 4, 5, 6, 1, 9}
	//stage 1 move dados da slice para o canal
	dataCh := sliceToChannel(n)
	//stage 2 data transformation (processing)
	productCh := sq(dataCh)
	//stage 3 consume data //end pipeline
	for n := range productCh {
		fmt.Println((n))
	} //print vai ser em ordem, pq a comunicação é sincronizada entre as rotinas.

}

func sliceToChannel(n []int) <-chan int {

	out := make(chan int) //criamos unbuffered channel pq?  // pois cada dado vai ser comunidade 1 por 1 em sincronia
	go func() {
		for _, n := range n {
			out <- n
		}
		close(out) //serve para avisar quando esta rotina finalizar para que a função leitora também possa fechar sua rotina
	}()
	return out //quando o ch recebe 1 valor o ch é retornado pra fora da função e a go routine segue executando
}

func sq(in <-chan int) <-chan int {

	out := make(chan int) //quando este canal lê a go rotine ele libera pra função de transformação 1
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out) // qdo recebe o valor de closed da função escritora, também finaliza sua rotina
	}()
	return out //retorn dado antes de fechar a rotina, printado na tela

}
