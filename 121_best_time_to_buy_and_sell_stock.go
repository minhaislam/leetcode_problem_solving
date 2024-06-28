package main

func main() {
	// val := 2
	// nums := []int{1, 1, 2, 3}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// val := 0
	nums := []int{7, 6, 4, 3, 1}

	// output := []int{1, 2, 2, 3, 5, 6}

	maxProfit(nums)

}

func maxProfit(prices []int) int {

	max_profit := 0
	// current_profit := 0
	min_value := prices[0]
	i := 1
	for len(prices) > 1 && i < len(prices) {
		// fmt.Println("Value for: ", min_value, prices[i])
		if min_value < prices[i] && max_profit < prices[i]-min_value {

			max_profit = prices[i] - min_value
			// min_value = prices[i]
		} else if min_value > prices[i] {
			min_value = prices[i]
			// max_profit = current_profit
		}
		// current_profit = prices[i] - min_value
		// if current_profit > max_profit {
		// 	max_profit = current_profit
		// }
		i++
	}
	// fmt.Println(max_profit)
	return max_profit

}
