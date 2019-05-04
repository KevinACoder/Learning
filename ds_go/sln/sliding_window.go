package sln

import "math"

//209. Minimum Size Subarray Sum
func minSubArrayLen(sum int, nums []int) (ans int) {
	var start, end, tmpSum int
	ans = math.MaxInt32
	for start < len(nums) {
		//move 'end' to make slide window valid
		for end < len(nums) && tmpSum < sum {
			tmpSum += nums[end]
			end++
		}
		if tmpSum < sum {
			break
		} //not able to find a valid sliding window
		ans = Min(ans, end-start)
		//move slide window by increase 'start'
		tmpSum -= nums[start]
		start++
	}

	if ans == math.MaxInt32 {
		ans = 0
	}
	return
}

//1004. Max Consecutive Ones III
func longestOnes(ones []int, k int) (ans int) {
	var start, end, zeros int
	for ; end < len(ones); end++ {
		if ones[end] == 0 {
			zeros++
		}
		//move 'start' util zeros are equal or less than k
		for zeros > k {
			//remove the head element, if it is zeros, reduce
			// num of zero in sliding window
			if ones[start] == 0 {
				zeros--
			}
			start++
		}
		ans = Max(ans, end-start+1)
	}
	return
}
