package om

import "fmt"

func Om() {
	var (
		//Operasi Matematika
		a = 10
		b = 2
		c = a + b
		d = a - b 
		e = a * b
		f = a / b
		g = a % b
		i =10
	)
	//Hasil
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	i += 10 // i + 10
	i++ // i = i + 1
	fmt.Println(i)
}