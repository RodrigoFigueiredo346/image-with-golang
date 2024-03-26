package examples

import "fmt"

func MaxProfit(p []int) int {

	purchase := 0
	sell := 0
	profit := 0
	lastSell := 0
	aux := 0
	maior := 0

	for i, v := range p { // 7,6,8,3,5,9,1,5
		if i == 0 {
			purchase = v
		} else {
			if v < purchase {

				for x := i; i < len(p); x++ {
					if p[x] > maior {
						maior = p[x]
						//profit = p[x] - purchase
					}
				}
				purchase = v
				profit = maior - purchase

			} else {
				if sell < v {
					sell = v
				}
			}

			profit = sell - purchase
		}
	}
	fmt.Println(aux, sell)

	if lastSell == 0 || purchase == 0 {
		return 0
	} else {
		return profit
	}

}
