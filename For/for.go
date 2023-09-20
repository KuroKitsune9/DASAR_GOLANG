package untuk

import "fmt"

func For() {
    

	for counter := 1;counter <= 10; counter++ {
		fmt.Println("ini pengualangan yang ke = ", counter)
	}

	slice := []string{"Muharafi", "Dalilah"}

	for i := 0; i < len(slice); i++ { 
	fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i , "Value = ", value)
	}

	person := make(map[string]string)
	person["name"] = "Muharafi"
	person["title"] = "Programmer"


	for key, value := range person { 
		fmt.Println(key, "=" , value)
	}

}