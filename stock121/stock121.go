package stock121

// import "fmt"

/*####  Algorithm

Input [7,1,5,3,6,4] => Output 5
1. Initiate Profit and MinStockPrice
	a.  Profit =0
	b.  MaxStockPrice = prices[0]
2. Iterate through the array from 1 as i
	CurrentProfit = prices[i]- MinStockPrice

	if CurrentProfit > Profit ##Get the large profit always
		Profit=CurrentProfit

	if Price[i]<MinStockPrice ##Get the small pick point always
		MinStockPrice=Price[i]

3. return Profit

				Input [2,4,1,7]  Output 6
				Input [2,7,1,4]	 Output 5
*/
func Profit(prices []int) int {

	Profit := 0
	MinStockPrice := prices[0]

	for i := 1; i < len(prices); i++ {

		CurrentProfit := prices[i] - MinStockPrice

		if CurrentProfit > Profit {
			Profit = CurrentProfit
		}
		if prices[i] < MinStockPrice {
			MinStockPrice = prices[i]
		}
	}
	return Profit
}
