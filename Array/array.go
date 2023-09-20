package array

import "fmt"

func Array() {
	var names [2]string
	names[0] = "Muharafi"
	names[1] = "Dalilah"

	fmt.Println(names[0])
	fmt.Println(names[1])

	//membuat array secara langsung
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(len(values))
}