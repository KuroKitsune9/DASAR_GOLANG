package sprint

import "fmt"

func Sprint() {
    quantity := 10
	fruits := "Apple"

	//Sprintf
	text := fmt.Sprint("I have ", quantity, " ", fruits)
	fmt.Printf("Value stored in text is: '%v' \n", text)
}