package Struk

import "fmt"
type Customer struct {
	Name, Address string
	age int
}


func Struct() {
    var Kuro Customer
	Kuro.Name = "KuroKitsune"
	Kuro.Address = "Jakarta"
	Kuro.age = 17

	fmt.Println(Kuro.Name)
	fmt.Println(Kuro.Address)
	fmt.Println(Kuro.age)
}