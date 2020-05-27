package main

import "fmt"

func main() {
	arr :=[]int  {3,2,4}
	fmt.Println(twoSum01(arr,6))
	var l104 =  ListNode{Val: 2, Next: nil}
	var l103 =  ListNode{Val: 3, Next: &l104}
	var l102 =  ListNode{Val: 4, Next: &l103}
	var l101 =  ListNode{Val: 1, Next: &l102}
	var l204 =  ListNode{Val: 3, Next: nil}
	var l203 =  ListNode{Val: 4, Next: &l204}
	var l202 =  ListNode{Val: 9, Next: &l203}
	var l201 =  ListNode{Val: 9, Next: &l202}
	scanLinklist(addTwoNumbers(&l101,&l201))
	fmt.Println(lengthOfLongestSubstring("abdwiajfwaljwalgawfegsvg"))
	fmt.Println(lengthOfLongestSubstring("abcabcaa"))

}
