package subarrary

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
