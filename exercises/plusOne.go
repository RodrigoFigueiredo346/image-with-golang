package exercises

import (
	"fmt"
	"strconv"
)

func PlusOne(digits []int) []int {
	val := ""
	for _, v := range digits {
		val = val + strconv.Itoa(v)
	}

	// Convertendo a string para um número int64 e incrementando
	num, _ := strconv.ParseInt(val, 10, 64)
	num++

	// Convertendo o número de volta para uma string com zeros à esquerda
	numStr := fmt.Sprintf("%0*d", len(digits), num)

	// Convertendo cada dígito da string de volta para um int e armazenando no slice resultante
	res := make([]int, len(numStr))
	for i, s := range numStr {
		aux, _ := strconv.Atoi(string(s))
		res[i] = aux
	}

	return res

}
