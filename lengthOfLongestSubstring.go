package main

//Longest Substring Without Repeating Characters
//Medium
//Given a string, find the length of the longest substring without repeating characters.
//Example 1:
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	var max = 0
	//var lenS = len(s)
	//var arr = strings.Split(s,"")
	//for index,cur :=range arr{
	//	//var posi = getLatestPosi(arr[index:],arr[index])
	//	//if(posi!=-1){
	//	//
	//	//}
	//}
	return max
}
//func getMaxSub(arr []rune , seek rune,count int) int{
//	var max = 0
//	for index,cur :=range arr{
//		getMaxSub(arr[])
//	}
//	return max
//}
//func testContains(arr []rune,seek rune) bool{
//	for index,cur :=range arr{
//		if(arr[index]==seek){
//			return true
//		}
//	}
//	return false
//
//}
//func getLatestPosi(arr []rune,seek rune) int{
//	for index,cur :=range arr{
//		if(arr[index]==seek){
//			return index
//		}
//	}
//	return -1
//}