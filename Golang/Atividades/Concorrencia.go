package main

import (
	"fmt"
	"sync"
	"time"
)

// Implemente um sistema onde:

// 10 tarefas são processadas por 3 workers simultaneamente -- FEITO
// O sistema deve cancelar todas as operações após 2 segundos -- FEITO
// Cada worker deve:
// Receber um ID de tarefa -- FEITO
// "Processar" (simular com time.Sleep) -- FEITO
// Reportar conclusão -- FEITO
// Use WaitGroup para esperar todos os workers finalizarem -- FEITO
// Se o timeout ocorrer, cancele as tarefas pendentes -- FEITO

func Worker(id int, jobs <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-done:
			fmt.Println("Sinal parar...")
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Acabaram as tarefas")
				return
			}
			fmt.Println(id, "Inicio", job)
			time.Sleep(1 * time.Second)
			fmt.Println(id, "Término", job)
		}
	}

	fmt.Println("")
}

func main() {
	const numWorkers = 3
	const numJobs = 10
	var wg sync.WaitGroup
	jobs := make(chan int)
	done := make(chan struct{})

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, jobs, done, &wg)
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	wg.Wait()
	fmt.Println("Rolê concluído")
}
