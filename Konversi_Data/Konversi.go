package Konversi

import "fmt"

func Konver() {
	var (
		nilai32 int32 = 100000;
		nilai64 int64 = int64(nilai32);
		nilai16 int16 = int16(nilai32);
		nilai8 int8 = int8(nilai32)
		//konversi dai byte ke string
		name = "Muharafi Dalilah"
		e = name[7]
		eString string = string(e) 
	)

	fmt.Println(nilai32);
	fmt.Println(nilai64);
	fmt.Println(nilai16);
	fmt.Println(nilai8);
	//hasil konversi byte ke string
	fmt.Println(name)
	fmt.Println(eString)
}