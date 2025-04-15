package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// defer says that whatever comes after this, dont execute it until the current function exists
	// e.g. wg.Done() will not be executed till content in this function is executed.
	//which means wg.Done is the last thing that will be executed in this function
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	// created a wait group
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	// wg.Add(9) -> there is a better alternative to write this (see next line)
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d : %s", i, x), &wg)
		// fmt.Sprintf("%d : %s", i, x) -> 8 : epsilon
	}

	wg.Wait()

	// below code is to avoid "negative WaitGroup counter" error
	// printSomething function is called at line no.50, there is a wg.Done() in it
	// we need to increament in wait group value only then wg.Done() will decrement it
	// before this line the wg count will be 0 and after calling printsomething -> wg.Done() it will become -1, which is not possible (for waitgroup to hold -ve value)
	wg.Add(1)

	// go printSomething("nothing to PRINT")
	// time.Sleep(1 * time.Second)
	printSomething("SOMEthing to PRint", &wg)
}
