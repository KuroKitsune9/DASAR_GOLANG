package variable

import "fmt"

func Var() {
	//multiple variable 
	var (
		FirstName = "Muharafi"
		LastName = "Dalilah"
	)

	fmt.Println(FirstName, LastName)

	//beragam variable di golang
	var name string
	var number int
	var nama = "Muharafi Dalilah"
	country := "Indonesia"

	name = "Muharafi  Dalilah"
	fmt.Println(name)

	name = "Sandy Saputra"
	fmt.Println(name)

	number = 12345
	fmt.Println(number)

	number = 678910
	fmt.Println(number)

	fmt.Println(nama)
	fmt.Println(country)
}