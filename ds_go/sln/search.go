package sln

import (
	"math"
	"sort"
)

//1014. Capacity To Ship Packages Within D Days
func shipWithinDays(weights []int, D int) int {
	maxEle := func(xs []int) (max int) {
		max = math.MinInt32
		for _, x := range xs {
			if x > max {
				max = x
			}
		}
		return
	}

	sumEle := func(xs []int) (sum int) {
		for _, x := range xs {
			sum += x
		}
		return
	}

	//the target daily capcity must be within [low, high]
	low := maxEle(weights)
	high := sumEle(weights) + 1
	for low < high {
		mid := low + (high-low)/2
		days := 1  //count days take to ship
		trans := 0 //count weights can be ship in one day
		//try to ship out all weights with selected dayily capcity
		//  mid
		for _, w := range weights {
			if trans+w > mid {
				days++ //need to ship in the next day
				trans += w
				trans = w
			} else {
				trans += w //can be ship today
			}
		}

		if days > D {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

//74. Search a 2D Matrix
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	if col == 0 {
		return false
	}

	low, high := 0, row*col
	//treat matrix as array and parse row/col index when need to
	//  access element
	for low < high {
		mid := low + (high-low)/2
		r, c := mid/col, mid%col
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}

//981. Time Based Key-Value Store
type TimeMap struct {
	kv map[string]map[int]string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	var tm TimeMap
	tm.kv = make(map[string]map[int]string)
	return tm
}

func (tm *TimeMap) Set(key string, val string, time int) {
	if _, found := tm.kv[key]; !found {
		tm.kv[key] = make(map[int]string)
	}
	tm.kv[key][time] = val
}

func (tm *TimeMap) Get(key string, time int) string {
	valSet, found := tm.kv[key]
	if !found || len(valSet) == 0 {
		return ""
	}
	var times []int
	for k := range valSet {
		times = append(times, k)
	}
	sort.Ints(times)
	idx := len(times) - 1
	//find the first time stamp before time
	for i := len(times) - 1; i > 0; i-- {
		if times[i] > time && times[i-1] < time {
			idx = i - 1
		}
	}
	val := valSet[times[idx]]
	return val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

//962. Maximum Width Ramp
func maxWidthRamp(nums []int) (ans int) {
	var stack []int
	top := func(s []int) int {
		if len(s) == 0 {
			panic("access out of boundary")
		}
		return s[len(s)-1]
	}
	pop := func(s []int) []int {
		if len(s) == 0 {
			panic("access out of boundary")
		}
		return s[:len(s)-1]
	}
	push := func(s []int, n int) []int {
		return append(s, n)
	}

	//push the start candiator
	for i, n := range nums {
		if len(stack) == 0 || n < nums[top(stack)] {
			//if current 'n' is greater than stack top
			//  it is impossible to be a candiator
			stack = push(stack, i)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		//find the end candiator for each of the stack top
		for len(stack) > 0 && nums[i] >= nums[top(stack)] {
			ans = Max(ans, i-top(stack))
			stack = pop(stack)
		}
	}
	return
}

//18. 4Sum
func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				t := target - nums[i] - nums[j] - nums[k]
				if t < nums[k] {
					break
				}
				if idx := sort.SearchInts(nums[k+1:], t) + k + 1; idx < len(nums) && nums[idx] == t {
					ans = append(ans, []int{nums[i], nums[j], nums[k],
						nums[idx]})
				}
			}
		}
	}
	return
}

//668. Kth Smallest Number in Multiplication Table
func findKthNumber(row, col, target int) int {
	LEX := func(x int) (cnt int) {
		for i := 1; i <= row; i++ {
			if x/i > col {
				cnt += col //count the whole row
			} else {
				cnt += x / i //count the ele before x, first several cols
			}
		}
		return
	}

	low, high := 1, row*col+1
	for low < high {
		mid := low + (high-low)/2
		cnt := LEX(mid)

		if cnt >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
