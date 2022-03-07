# DailyCodingProblem513

## 問題敘述

給定一列整數還有一個數字 K ， 回傳數列中包含連續元素的子數列，子數列的和等於 K 。

舉例來說，如果給定數列 [1, 2, 3, 4, 5] 且 K = 9 ， 則回傳 [2, 3, 4] ， 因為 2 + 3 + 4 = 9 。

## 我的解法

要找出特定值 K 和的子陣列

需要符合兩個條件

1 所有元素都是鄰近的

2 和為特定值

因此可以, 從當下位址的值往連續加到某一個值接近直到和的等於 K 或是走到陣列尾部

如果加到某一個位址的值剛好是該陣列的值

則回傳該子陣列

如果加到陣列尾部都不等於 K 則繼續從下一個位址加


## 程式碼

```golang
func SubArraySum(nums []int, K int) []int {
	var result []int
	total := len(nums)
	start := 0
	for i := range nums {
		sum := 0
		start = i
		for last := i; last < total; last++ {
			if sum+nums[last] == K {
				result = make([]int, last-start+1)
				copy(result, nums[start:last+1])
				return result
			}
			sum += nums[last]
		}
	}
	return result
}
```

