package lanjut

import "fmt"

func Lanjut() {
    for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		fmt.Println("Perulangan ke =" ,i)
	}

}