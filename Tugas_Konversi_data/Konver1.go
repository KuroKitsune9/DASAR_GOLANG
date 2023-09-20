package tugaskonversidata

import (
	"fmt"
	"strconv"

)

func Konversi() {
	// int32 to int64
	var int32Value = int32(123)
	int64Value := int64(int32Value)
	fmt.Println(int64Value)
	// int64 to int32
	int64Value = int64(123)
	int32Value = int32(int64Value)
	fmt.Println(int32Value)

	// string to byte
	str := "Hello"
	char := str[0]
	fmt.Printf("character %c\n", char)
	
	//byte to string
	str = strconv.Itoa(int(char))
    fmt.Printf("string %s\n", str)

	//string to rune
	str2 := "Hello"
	runeVal := []rune(str2)
	fmt.Println(runeVal)

	//rune to string
	str9:= string(runeVal)
    fmt.Println(str9)

	//int64 to string
	int64Value3 := int64(189)
	str3 := strconv.FormatInt(int64Value3, 10)
	fmt.Println(str3)

	//int32 to string
	int32Value = int32(586)
	str4 := strconv.Itoa(int(int32Value))
	fmt.Println(str4)

	//string to int64
	str5 := "456"
	int64Value4, _ := strconv.ParseInt(str5, 10, 64)
	fmt.Println(int64Value4)

	//string to int32
	str6 := "666"
	int32Value4, _ := strconv.Atoi(str6)
	fmt.Println(int32Value4)

	//string  to bool
	str7 := "true"
    boolVal, _ := strconv.ParseBool(str7)
    fmt.Println(boolVal)

    //bool to string
    boolVal2 := true
    str8 := strconv.FormatBool(boolVal2)
    fmt.Println(str8)
}
