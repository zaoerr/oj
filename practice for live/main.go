package main

import (
	"fmt"
	"math"
)

func Solution(arr []int, k int) int {
	lp, rp, sum := 0, 0, 0
	mL := math.MaxInt
	for rp < len(arr) {
		sum += arr[rp]
		for sum >= k && lp <= rp {
			mL = min(mL, rp-lp+1)
			sum -= arr[lp]
			lp++
		}
		rp++
	}
	return mL
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(Solution(arr, 5))
	return
}
