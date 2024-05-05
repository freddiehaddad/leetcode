package arrays

// You are given an integer array nums. You are initially positioned at the
// array's first index, and each element in the array represents your maximum
// jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
func canJump(nums []int) bool {
	var position int
	for i := 1; i < len(nums); i++ {
		remaining := nums[position] - (i - position)
		if remaining < 0 {
			return false
		}

		if nums[i] > remaining {
			position = i
		}
	}
	return true
}

// You are given a 0-indexed array of integers nums of length n. You are
// initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward jump from
// index i. In other words, if you are at nums[i], you can jump to any
// nums[i+j] where:
//
//   1. 0 <= j <= nums[i]
//   2. i + j < n
//
// Return the minimum number of jumps to reach nums[n-1]. The test cases are
// generated such that you can reach nums[n-1].
//
// Constraints:
//
//   1. 1 <= nums.length <= 10^4
//   2. 0 <= nums[i] <= 1000
//   3. It's guaranteed that you can reach nums[n-1]
func canJumpII(nums []int) int {
	var farthest, jumps, nextJump int
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == nextJump {
			nextJump = farthest
			jumps += 1
		}
	}
	return jumps
}

// There are n gas stations along a circular route, where the amount of gas at
// the ith station is gas[i].
//
// You have a car with an unlimited gas tank and it costs cost[i] of gas to
// travel from the ith station to its next (i + 1)th station. You begin the
// journey with an empty tank at one of the gas stations.
//
// Given two integer arrays gas and cost, return the starting gas station's
// index if you can travel around the circuit once in the clockwise direction,
// otherwise return -1. If there exists a solution, it is guaranteed to be
// unique
func canCompleteCircuit(gas []int, cost []int) int {
	var currentPosition, startingPosition, tank, nPositions int
	nPositions = len(gas)
	tank = gas[0]
	for {
		nextPosition := (currentPosition + 1) % nPositions
		if cost[currentPosition] > tank {
			if nextPosition <= startingPosition {
				startingPosition = -1
				break
			}

			currentPosition = nextPosition
			startingPosition = nextPosition
			tank = gas[currentPosition]
			continue
		}

		tank -= cost[currentPosition]
		currentPosition = nextPosition
		tank += gas[currentPosition]

		if currentPosition == startingPosition {
			break
		}
	}

	return startingPosition
}

// There are n children standing in a line. Each child is assigned a rating
// value given in the integer array ratings.
//
// You are giving candies to these children subjected to the following
// requirements:
//
//   1. Each child must have at least one candy.
//   2. Children with a higher rating get more candies than their neighbors.
//
// Return the minimum number of candies you need to have to distribute the
// candies to the children.
func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	// distribute candies from left to right
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = max(candies[i], candies[i-1]+1)
		}
	}

	// distribute candies from right to left
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// total up the candies
	var sum int
	for _, c := range candies {
		sum += c
	}

	return sum
}

// Given an array of integers citations where citations[i] is the number of
// citations a researcher received for their ith paper, return the researcher's
// h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined
// as the maximum value of h such that the given researcher has published at
// least h papers that have each been cited at least h times.
func hIndex(citations []int) int {
	n := len(citations)
	counts := make([]int, n+1)
	for _, c := range citations {
		if c > n {
			counts[n]++
		} else {
			counts[c]++
		}
	}
	var count int
	for ; n > 0; n-- {
		count += counts[n]
		if count >= n {
			return n
		}
	}
	return 0
}

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

// You are given an integer array prices where prices[i] is the price of a
// given stock on the ith day.
//
// On each day, you may decide to buy and/or sell the stock. You can only hold
// at most one share of the stock at any time. However, you can buy it then
// immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
func maxProfitII(prices []int) int {
	price := prices[0]
	profit := 0

	for _, val := range prices[1:] {
		if p := val - price; p > 0 {
			profit += p
			price = val
		}

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

// Given an integer array nums, return an array answer such that answer[i] is
// equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
// integer.
//
// You must write an algorithm that runs in O(n) time and without using the
// division operation.
func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	var multiply func(i, leftProduct int) int
	multiply = func(i, leftProduct int) int {
		// base case
		if i == len(nums)-1 {
			products[i] = leftProduct
			return nums[i]
		}
		result := multiply(i+1, leftProduct*nums[i])
		products[i] = leftProduct * result
		return result * nums[i]
	}
	products[0] = multiply(1, nums[0])
	return products
}

// Given an integer array nums sorted in non-decreasing order, remove the
// duplicates in-place such that each unique element appears only once. The
// relative order of the elements should be kept the same. Then return the
// number of unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted, you
// need to do the following things:
//
//   1. Change the array nums such that the first k elements of nums contain
//      the unique elements in the order they were present in nums initially.
//      The remaining elements of nums are not important as well as the size of
//      nums.
//   2. Return k.
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
//   1. Change the array nums such that the first k elements of nums contain
//      the elements which are not equal to val. The remaining elements of nums
//      are not important as well as the size of nums.
//   2. Return k.
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
