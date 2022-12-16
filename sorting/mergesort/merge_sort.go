package mergesort

func MergeSort(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, startIndex, endIndex int) []int {
	if endIndex-startIndex <= 0 {
		return []int{nums[startIndex]}
	}

	middle := len(nums) / 2
	lRef := nums[:middle]
	rRef := nums[middle:]

	lSide := make([]int, len(lRef))
	rSide := make([]int, len(rRef))

	copy(lSide, lRef)
	copy(rSide, rRef)

	lRes := sort(lSide, 0, len(lSide)-1)
	rRes := sort(rSide, 0, len(rSide)-1)

	return merge(lRes, rRes)
}

func merge(lSides, rightSides []int) []int {
	results := make([]int, len(lSides)+len(rightSides))

	var lWalker int
	var rWalker int

	for i := 0; i < len(results); i++ {
		if lWalker > len(lSides)-1 {
			results[i] = rightSides[rWalker]
			rWalker++
			continue
		} else if rWalker > len(rightSides)-1 {
			results[i] = lSides[lWalker]
			lWalker++
			continue
		}

		if lSides[lWalker] <= rightSides[rWalker] {
			results[i] = lSides[lWalker]
			lWalker++
		} else {
			results[i] = rightSides[rWalker]
			rWalker++
		}
	}

	return results
}
