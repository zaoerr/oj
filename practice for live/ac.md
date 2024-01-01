## k优雅阈值
### 如果一个数组中的某个子段,其中出现次数最多的元素出现大于等于 x 次， 被称为x−子段
### hint
双指针，可作为模板
### code
```
package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	cnt := make(map[int]int, 0)
	ans, l, r := 0, 0, 0
	for r < n {
		cnt[arr[r]]++
		for cnt[arr[r]] >= k {
			ans += n - r
			cnt[arr[l]]--
			l++
		}
		r++
	}
	fmt.Println(ans)
	return
}
```