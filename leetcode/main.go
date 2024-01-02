package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	type pair struct {
		cnt1 int
		cnt2 int
	}
	recall := make(map[int]pair, 0)
	pre, cur := new(pair), new(pair)
	cnt1, cnt2, idx := 0, 0, 0
	for true {
		cnt1++
		for _, c := range s1 {
			if int(c) == int(s2[idx]) {
				idx++
				if idx == len(s2) {
					cnt2++
					idx = 0
				}
			}
		}

		if cnt1 == n1 {
			return cnt2 / n2
		}

		if v, ok := recall[idx]; ok {
			pre.cnt1, pre.cnt2 = v.cnt1, v.cnt2
			cur.cnt1, cur.cnt2 = cnt1-v.cnt1, cnt2-v.cnt2
			break
		} else {
			recall[idx] = pair{cnt1, cnt2}
		}
	}

	ans := pre.cnt2 + (n1-pre.cnt1)/cur.cnt1*cur.cnt2
	rest := (n1 - pre.cnt1) % cur.cnt1
	for i := 0; i < rest; i++ {
		for _, c := range s1 {
			if int(c) == int(s2[idx]) {
				idx++
				if idx == len(s2) {
					ans++
					idx = 0
				}
			}
		}
	}
	return ans / n2
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	dp := make([]int, n2)
	for i := 0; i < len(s2); i++ {
		j := i
		for k := 0; k < len(s1); k++ {
			if s2[j] == s1[k] {
				j = (j + 1) % len(s2)
				dp[i]++
			}
		}
	}
	idx, ans := 0, 0
	for i := 0; i < n1; i++ {
		delta := dp[idx]
		idx = (idx + delta) % len(s2)
		ans += delta
	}

	return ans / (len(s2) * n2)
}

func main() {

	return
}
