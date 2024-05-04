package arrays

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// See: Boyer-Moore majority vote algorithm
func majorityElement(nums []int) int {
	c := 1
	m := nums[0]
	for _, val := range nums[1:] {
		if c == 0 {
			m = val
			c = 1
		} else if m == val {
			c++
		} else {
			c--
		}
	}
	return m
}

// You are given an array prices where prices[i] is the price of a given stock
// on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you
// cannot achieve any profit, return 0.
func maxProfit(prices []int) int {
	price := prices[0]
	profit := 0

	for _, val := range prices[1:] {
		profit = max(profit, val-price)
		price = min(price, val)
	}
	return profit
}

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing
// order, and two integers m and n, representing the number of elements in
// nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead
// be stored inside the array nums1. To accommodate this, nums1 has a length of
// m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a
// length of n.
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}

	for n >= 0 {
		nums1[i] = nums2[n]
		i--
		n--
	}
}

// Given an integer array nums sorted in non-decreasing order, remove the
// duplicates in-place such that each unique element appears only once. The
// relative order of the elements should be kept the same. Then return the
// number of unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted, you
// need to do the following things:
//
//   - Change the array nums such that the first k elements of nums
//     contain the unique elements in the order they were present in
//     nums initially. The remaining elements of nums are not important
//     as well as the size of nums.
//   - Return k.
func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// Given an integer array nums sorted in non-decreasing order, remove some
// duplicates in-place such that each unique element appears at most twice. The
// relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array
// nums. More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does not
// matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying
// the input array in-place with O(1) extra memory.
func removeDuplicatesII(nums []int) int {
	k := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// Given an integer array nums and an integer val, remove all occurrences of
// val in nums in-place. The order of the elements may be changed. Then return
// the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to
// get accepted, you need to do the following things:
//
//   - Change the array nums such that the first k elements of nums
//     contain the elements which are not equal to val. The remaining
//     elements of nums are not important as well as the size of nums.
//   - Return k.
func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// Given an integer array nums, rotate the array to the right by k steps, where
// k is non-negative.
func rotate(nums []int, k int) {
	k = k % len(nums)

	// reverse entire array
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse first k elements
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse the last k elements
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
