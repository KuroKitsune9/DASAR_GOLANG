package kutipinstring

import (
	"fmt"
	"strconv"

)

func Kutip(){
	//print string with double quotes
	str := "Muharafi Dalilah"
	// fmt.Printf("%q/n" ,str)
	//cara ke dua
	fmt.Println(strconv.Quote(str))

	//create string with double quotes
	
	quoted := strconv.Quote(str)
	fmt.Println(quoted)
	//cara ke dua




}