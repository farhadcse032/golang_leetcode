package stock121

import "fmt"

/*####  Algorithm

Input [7,1,5,3,6,4] => Output 5
1. Initiate two pick point
	a.  MinStockPrice =prices[0]
	b.  MaxStockPrice =0
2. Iterate through the array from 1 as i
	if Price[i] < MinStockPrice
		MinStockPrice=Price[i]
	if Price[i]>MinStockPrice &&  Price[i]>MaxStockPrice
		MaxStockPrice=Price[i]
3. return MaxStockPrice-MinStockPrice if MaxStockPrice>MinStockPrice else 0

{2,1,2,0,1}
*/
func Profit(prices []int) int {
	MinStockPrice := prices[0]
	MaxStockPrice := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] == 0 {
			MinStockPrice = 0
			MaxStockPrice = 0

		} else if prices[i] < MinStockPrice && prices[i] > MaxStockPrice {
			fmt.Println(MinStockPrice, MaxStockPrice, prices[i])
			MinStockPrice = prices[i]
		}
		if prices[i] > MinStockPrice && prices[i] > MaxStockPrice {

			MaxStockPrice = prices[i]
		}
	}
	if MaxStockPrice > MinStockPrice {
		fmt.Println(MinStockPrice, MaxStockPrice)
		return MaxStockPrice - MinStockPrice
	}
	return 0
}
