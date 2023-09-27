package peta

import "fmt"

func Map(){
	person := map[string]string{
		"name": "Muharafi",
        "address": "Bandung",
		"phone": "08123456789",
	}
	person["title"] = "Progammers" 
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["phone"])

	var book map[string]string =  make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

	for k, v := range book {
		fmt.Println(v, k)
	}
}