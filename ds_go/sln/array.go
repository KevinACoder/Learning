package sln

//941. Valid Mountain Array
func validMountainArray(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var hasUp, hasDown bool
	//track the up and down
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			//fasle if up followed down
			if hasDown {
				return false
			}
			hasUp = true
		} else if nums[i] < nums[i-1] {
			//false if down comes before up
			if !hasUp {
				return false
			}
			hasDown = true
		} else {
			return false
		}
	}
	return hasUp && hasDown
}

//896. Monotonic Array
func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	isDec, isInc := true, true

	for i := 1; i < len(nums); i++ {
		isInc = isInc && nums[i] >= nums[i-1]
		isDec = isDec && nums[i] <= nums[i-1]
		//break if not monotonic
		if !isInc && !isDec {
			return false
		}
	}

	return isInc || isDec
}

func fairCandySwap(A, B []int) (ans []int) {
	ans = make([]int, 2)
	//calc sum and diff of two array
	sumA, sumB := Sum(A), Sum(B)
	diff := (sumA - sumB) / 2

	//build a set of A
	tbl := make(map[int]int)
	for _, a := range A {
		tbl[a]++
	}

	//check if exist a valid ans
	for _, b := range B {
		if tbl[b+diff] > 0 {
			ans[0] = b + diff
			ans[1] = b
		}
	}

	return
}
