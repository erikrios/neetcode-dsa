package bucketsort

func BucketSort(nums []int) {
	if len(nums) < 5 {
		return
	}

	var buckets [5]int

	for i := 0; i < len(nums); i++ {
		buckets[nums[i]]++
	}

	var counter int
	for i := 0; i < len(buckets); i++ {
		for j := 0; j < buckets[i]; j++ {
			nums[counter] = i
			counter++
		}
	}
}
