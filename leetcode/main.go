package main

import (
	"strconv"
	"strings"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(finish, 10)
	n := len(high)
	diff := n - len(s)
	low = strings.Repeat("0", n-len(low)) + low

	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int64
	dfs = func(i int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			return 1
		}
		// be careful!
		if !limitLow && !limitHigh {
			p := &memo[i]
			if *p != -1 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		if i < diff {
			for d := lo; d <= min(hi, limit); d++ {
				res += dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
			}
		} else {
			x := int(s[i-diff] - '0')
			if x >= lo && x <= min(hi, limit) {
				res += dfs(i+1, limitLow && x == lo, limitHigh && x == hi)
			}
		}
		return
	}
	return dfs(0, true, true)
}

func main() {
	return
}
