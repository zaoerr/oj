package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	ans, profits, maxProfits := -1, 0, 0
	for i, waits := 0, 0; waits > 0 || i < len(customers); i++ {
		if i < len(customers) {
			waits += customers[i]
		}
		up := min(4, waits)
		waits -= up
		profits += up*boardingCost - runningCost
		if profits > maxProfits {
			maxProfits = profits
			ans = i + 1
		}
	}
	return ans
}

func main() {

	return
}
