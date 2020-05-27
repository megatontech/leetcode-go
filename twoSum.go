package main
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum01(nums []int, target int) []int {
	var result []int = make([]int,2)
	for i, i2 := range nums {
		subnums := nums[i+1:]
		for j, j2 := range subnums {
			if(i2+j2==target){
				result[0]=i
				result[1]=j+i+1
			}
		}
	}
	return result
}
