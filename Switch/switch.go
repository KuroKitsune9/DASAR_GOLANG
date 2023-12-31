package ganti

import "fmt"

func Switch() {
	names := "Muharafi"

	switch {
	case names == "Muharafi":
		fmt.Println("hello Muharafi")
		fmt.Println("hello Muharafi")
	case names == "Muhammad":
		fmt.Println("hello Muhammad")
		fmt.Println("hello Muhammad")
	case len(names) > 3 && names != "Muharafi":
		fmt.Println("hello Muhammad")
	default: 
		fmt.Println("hello, boleh kenalan?")
		fmt.Println("hello, boleh kenalan?")
	}

	//switch dengan short statement
	// switch length := len(names); length > 5 {
	// 	case true:
    //         fmt.Println("Nama Terlalu Panjang")
    //     default:
    //         fmt.Println("Nama Terlalu Pendek")
	// }

	length := len(names)
	switch {
		case length > 10:
            fmt.Println("Nama Terlalu Panjang")
		case length > 5:
			fmt.Println("Nama Lumayan Panjang")
        default:
            fmt.Println("Nama Sudah Benar")
	}
}