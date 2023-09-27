package clorsure

import "fmt"

func Closure() {
	name := "muha"
    counter := 0


	increment := func() {
		name = "Dalil"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	println(counter)
	println(name)
}