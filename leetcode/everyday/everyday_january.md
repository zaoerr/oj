## 1599. 经营摩天轮的最大利润
### Title:
https://leetcode.cn/problems/maximum-profit-of-operating-a-centennial-wheel/
### Idea:
以等待的人 wait 和 顾客数组下标 i 为迭代变量，坐上摩天轮的游客数量 up = min(4, wait)
### Code:
```
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	ans, profits, maxProfits := -1, 0, 0
	for i, wait := 0, 0; wait > 0 || i < len(customers); i++ {
		if i < len(customers) {
			wait += customers[i]
		}
		up := min(4, wait)
		wait -= up
		profits += up*boardingCost - runningCost
		if profits > maxProfits {
			maxProfits = profits
			ans = i+1
		}
	}
	return ans
}
```

## 466. 统计重复个数
### Title:
https://leetcode.cn/problems/count-the-repetitions/description/

定义 str = [s, n] 表示 str 由 n 个字符串 s 连接构成。

例如，str == ["abc", 3] =="abcabcabc" 。
如果可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。

例如，根据定义，s1 = "abc" 可以从 s2 = "abdbec" 获得，仅需要删除加粗且用斜体标识的字符。
现在给你两个字符串 s1 和 s2 和两个整数 n1 和 n2 。由此构造得到两个字符串，其中 str1 = [s1, n1]、str2 = [s2, n2] 。

请你找出一个最大整数 m ，以满足 str = [str2, m] 可以从 str1 获得。
### Idea:
* 循环节：
核心是对于重复的串s1组成的模式串，与多个s1的最后一个字符匹配的s2的字符下标i重复出现，则从i-len(s2)+0-i这段被称为循环节。
很显然符合贪心，以及贪心是正确解
* 动态规划：dp[i] 代表从s2的下标i开始，在0-len(s1)中总共可以匹配多个少个字符，注意如果匹配之后i的增长，i=(i+1)%len(s2)。
### Code:
```
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
```

## 从链表中删除节点
### Title
给你一个链表的头节点 head 。
移除每个右侧有一个更大数值的节点。
返回修改后链表的头节点 head 。
### Idea
* 递归
* 翻转链表，找右侧小于，翻转链表
### Code
```Go
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	cur := removeNodes(head.Next)
	if cur.Val > head.Val {
		return cur
	}
	head.Next = cur
	return head
}*
```

```Go
func removeNodes(head *ListNode) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var pre, cur, nxt *ListNode = nil, head, nil
		for cur != nil {
			nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		return pre
	}
	head = reverse(head)
	cur := head
	for cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	head = reverse(head)
	return head
}
```