package functionparameter

import "fmt"

func sayHelloTo(FirstName string, LastName string){
	fmt.Println("Hello", FirstName, LastName)
}
func try(){
	sayHelloTo("Muharafi", "Dalilah")
}