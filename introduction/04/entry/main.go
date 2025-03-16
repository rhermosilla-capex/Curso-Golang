package main

import (
	"fmt"
	pokemon "pokemon-concurrency/internal"
	"pokemon-concurrency/internal/storage/pokeapi"
	"sync"
	"time"
)

func routine(processChannel chan func() pokemon.PokemonDetail, dataChannel chan pokemon.PokemonDetail) {
	for process := range processChannel {
		dataChannel <- process()
	}

}

func printResults(dataChannel chan pokemon.PokemonDetail, charmanderChannel chan pokemon.PokemonDetail) {
	for data := range dataChannel {
		// Verificar si es charmander para tratarlo especialmente
		if data.Name == "charmander" {
			charmanderChannel <- data
		}

	}

}

func printCharmander(charmanderChannel chan pokemon.PokemonDetail) {
	for data := range charmanderChannel {
		time.Sleep(1500 * time.Millisecond) // Pausa para destacarlo
		fmt.Println("¡ENCONTRADO CHARMANDER!", data.Name)
	}
}

func main() {
	processChannel := make(chan func() pokemon.PokemonDetail)
	dataChannel := make(chan pokemon.PokemonDetail)
	charmanderChannel := make(chan pokemon.PokemonDetail)

	repo := pokeapi.NewPokeapiRepository()
	pokemons, err := repo.GetPokemons()

	if err != nil {
		fmt.Println(err)
		return
	}

	wg := &sync.WaitGroup{}

	// Iniciar goroutines para procesar y mostrar resultados
	go routine(processChannel, dataChannel)
	go printResults(dataChannel, charmanderChannel)
	go printCharmander(charmanderChannel)

	// Enviar cada pokemon para procesar
	for _, p := range pokemons {
		wg.Add(1)
		pokemonName := p.Name // Capturar el valor actual de p.Name
		pokemonId := p.Id     // Capturar el ID actual

		processChannel <- func() pokemon.PokemonDetail {
			fmt.Println("Ejecutando proceso de :", pokemonName)
			defer wg.Done()
			result, err := repo.GetPokemonDetail(pokemonId)
			if err != nil {
				fmt.Println(err)
			}
			return result
		}
	}

	// Esperar a que todas las solicitudes terminen
	wg.Wait()

	// Cerrar canales y esperar a que terminen las goroutines
	close(processChannel)

	close(dataChannel)

	fmt.Println("Proceso completado")
}
