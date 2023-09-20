package berhenti

import "fmt"

func Break() {
    for i := 0; i <= 10; i++{
		if i == 5 {
			continue
		}
		fmt.Println("Perulangan Ke =" ,i)
		}
}