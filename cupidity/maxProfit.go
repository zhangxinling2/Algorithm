package main

func maxProfit(prices []int) int {
	res := 0
	prices = append(prices, prices[len(prices)-1]-1)
	prevDiff, curDiff := 0, 0
	low := 0
	for i := 1; i < len(prices); i++ {
		curDiff = prices[i] - prices[i-1]
		if prevDiff <= 0 && curDiff > 0 {
			prevDiff = curDiff
			low = i - 1
		}
		if prevDiff >= 0 && curDiff < 0 {
			prevDiff = curDiff
			res += prices[i-1] - prices[low]
		}
	}
	return res
}
