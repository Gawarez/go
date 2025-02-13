package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Aqui é setada uma seed aleatória usando nanosegundos com unix time
	// pra garantia que as frases sejam sempre escolhidas aleatoriamente
	// em cada execução do programa.
	rand.Seed(time.Now().UnixNano())

	// Slice de frases que serão printadas aleatoriamente
	messages := []string{
		"Hello, World!",
		"How's it going?",
		"Go is awesome!",
		"What's your favorite programming language?",
		"Have a great day!",
		"Keep coding!",
		"Embrace the randomness!",
	}

	// Loop infinito que pega frases do slice aleatoriamente, passando
	// como índice um número aleatório com limite do tamanho do slice.
	for {
		randomIndex := rand.Intn(len(messages))
		fmt.Println(messages[randomIndex])
		time.Sleep(500 * time.Millisecond)
	}
}
