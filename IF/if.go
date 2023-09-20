package jika

import "fmt"

func If() {
	var name = "Muharafi Dalilah"

	if name == "eko" {
        fmt.Println("hallo eko")
    } else if name == "joko"{
        fmt.Println("hallo joko")
    } else {
		fmt.Println("hallo, kenalan yuk")
	}

	// short statemen if else
	if length := len(name); length > 5 {
		fmt.Println("Nama Sudah Passssssssss")
	} else {
		fmt.Println("Nama Belum Passssssssss")
	}
}