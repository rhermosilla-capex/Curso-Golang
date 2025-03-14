package main

import (
	selectexample "concurrency/tutorial/select/select"
	"concurrency/tutorial/select/waitgroups"
	"fmt"
	"sync"
	"time"
)

func selectExample() {
	a := make(chan string)
	b := make(chan string)
	go selectexample.Routine(&a, &b)
	a <- "hello"
	b <- "world"
	a <- "hello"
	a <- "hello"

	b <- "world"

	fmt.Println(<-a)

}

func printa() string {
	return "a"
}

func printb() string {
	return "b"
}

// waitgroupsExample demuestra el uso de WaitGroups en Go para sincronización de goroutines.
/*
 * Explicación de WaitGroup en Go:
 *
 * Un WaitGroup es una estructura de sincronización que permite a una goroutine
 * esperar a que un conjunto de goroutines finalice su ejecución antes de continuar.
 *
 * Métodos principales:
 * - Add(n): Incrementa el contador interno por n, indicando que hay n goroutines que rastrear
 * - Done(): Decrementa el contador interno, indicando que una goroutine ha completado su trabajo
 * - Wait(): Bloquea hasta que el contador interno llegue a 0, es decir, todas las goroutines han terminado
 */
func waitgroupsExample() {
	// Creamos dos canales para comunicación entre goroutines
	a := make(chan func() string)
	b := make(chan func() string)

	// Inicializamos un WaitGroup para sincronizar la goroutine principal con WGRoutine
	wg := &sync.WaitGroup{}

	// Incrementamos el contador interno del WaitGroup en 1
	// Esto indica que tenemos 1 goroutine que debe completarse antes de continuar
	wg.Add(2)

	// Lanzamos la goroutine que se ejecutará en paralelo
	// Le pasamos el WaitGroup para que pueda indicar cuando ha terminado
	go waitgroups.WGRoutine(a, b, wg)

	// Enviamos valores al canal 'a' que serán procesados por la goroutine
	a <- printa
	time.Sleep(1 * time.Second)
	b <- printb

	// Llamamos a Done() para decrementar el contador del WaitGroup
	// En este caso, esto simula la finalización del trabajo principal

	// Llamamos a Wait() para bloquear la ejecución hasta que todas las goroutines
	// indiquen que han finalizado (cuando el contador llegue a 0)
	wg.Wait() // Importante: esto faltaba en el código original

}

func main() {

	// selectExample()
	waitgroupsExample()

}
