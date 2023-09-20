package slice

import "fmt"

func Slice() {
	var months = [...]string{
		"January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
}

	var slices1 = months[4:7]
	fmt.Println(slices1)
	fmt.Println(len(slices1))
	fmt.Println(cap(slices1))

	// months[5] = "Diubah"
	// fmt.Println(slices1)

	// slices1[0] = "mei lagi"
	// fmt.Println(months)

	var slices2 = months[10:]
	fmt.Println(slices2)

	var slices3 = append(slices2, "Muharafi")
	fmt.Println(slices3)
	slices3[1] = "Bukan Desember"
	fmt.Println(slices3)
	fmt.Println(slices2)
	fmt.Println(months)

	newSlices := make([]string, 2, 5)
	
	newSlices[0] = "Muharafi"
	newSlices[1] = "Dalilah"

	fmt.Println(newSlices)

	copySlice := make([]string, 1, cap(newSlices))
	copy(copySlice, newSlices)
	fmt.Println(copySlice)

	// array sudah di tentukan si isi arraynya dan untuk slice dia di gunakan untuk memotong suatu array
	iniArray := [5] int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}