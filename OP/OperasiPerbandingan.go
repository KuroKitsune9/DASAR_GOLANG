package op

import "fmt"

func Op() {
	var (
		Value1       = 100
		Value2       = 200

		result1 bool = Value1 > Value2
		result2 bool = Value1 < Value2
		result3 bool = Value1 == Value2
		result4 bool = Value1 != Value2
	)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}