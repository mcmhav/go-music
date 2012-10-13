package main

import (
	"fmt"
)

type Cake struct {
	Size  int
	Color string
}

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func New() *Cake {
	var c *Cake
	c.Size = 2
	return c
}

func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	c := New()
	fmt.Println(c)
}

//Jeg synes du skal laste ned musikk av disse:

// - David Bowie (en eller annen "best of"-plate)
// - John Cale (Fragments of a raint season, live)
// - Beatles ("1", "Sgt. Pepper|s Lonely Hearts Club Band" og "Love")
// - Thåström (Skebokvarnsv. 209)
// - The Kooks (cd-er med sangene "Oog la" og "She moves in her own way")
// - Velvet underground ("The Velvet underground & Nico")
// - Lou Reed (Transformer)
// - Cornelis Vreeswijk (Guldkorn från mäster cees memoarer)
// - Jokke med tourettes/valentinerne ("Billig lykke", "Prisen for popen", "Trygge oslo")
// - 
