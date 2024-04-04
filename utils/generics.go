package utils

import "fmt"

func Generics() {
	sum([]int32{1, 2, 3, 4, 5})

	sum([]float32{1.7, 2.9, 8.8})
}

func sum[G int32 | float32](params []G) G {
	var ret G
	for _, value := range params {
		ret += value
	}

	fmt.Println(ret)

	return ret
}
