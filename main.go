package main

import (
	orms "main/orms/sqlc"
)

func main() {

	// param := []string{
	// 	"kjrqmzv9mmtxhgvsevenhvq7",
	// 	"four2tszbgmxpbvninebxns6nineqbqzgjpmpqr",
	// 	"rkzlnmzgnk91zckqprrptnthreefourtwo",
	// 	"fouronevzkbnzm6seven47",
	// }

	//res := examples.LongestCommonPrefix([]string{"denti", "denti", "casarao", "dente"})
	//res := examples.IsValid("{(())[[(]{}}")
	//res := exercises.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 8, 9, 9})
	//res := examples.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	//res := examples.StrStr("sadbutsad", "sad")
	//res := examples.PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	//res := exercises.MySqrt(415)
	//res := examples.SumDigits2()
	//res := examples.MaxProfit([]int{2, 4, 1})

	//CreateJSON()
	//utils.TestPostgreGorm()
	orms.TestPostgreSqlc()

	//fmt.Println(res)

}
