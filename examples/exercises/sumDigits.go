package examples

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func SumDigits() int {

	file := "SumDigits.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v", err)
		return 0
	}

	rows := strings.Split(string(content), "\n")

	//elimina última linha em branco
	if rows[len(rows)-1] == "" {
		rows = rows[:len(rows)-1]
	}

	res := 0
	listNum := []int{}
	aux := ""
	var digit string

	for _, v := range rows {
		digit = ""
		for ind := 0; ind < len(v); ind++ {
			if unicode.IsDigit(rune(v[ind])) {

				digit += string(v[ind])
			}
		}
		if digit != "" {
			aux = string(digit[0]) + string(digit[len(digit)-1])

		} else {
			fmt.Println("sem dígitos...")
			continue
		}
		aux2, _ := strconv.Atoi(aux)
		listNum = append(listNum, aux2)

	}

	for _, val := range listNum {

		res += val
	}
	return res
}
