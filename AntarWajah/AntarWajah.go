package antarwajah

import "fmt"

type Hasname interface {
	GetName() string
}

func antarwajah(hasName Hasname){
	fmt.Printf("Hello", hasName.GetName())
}

func antarwajah2(){
	
}