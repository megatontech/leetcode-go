package main

import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func scanLinklist(p *ListNode){     // 形参p是Student型的指针，指针参数需要传入一个地址
	for p != nil {
		fmt.Println(*p)
		p = p.Next      // 这是简写，标准写法是   p = (*p).next
	}
}
//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example:
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 =  ListNode{Val: 0, Next: nil}
	var offset,count = 0,0
	var lastNode = &l3
	for (l1!=nil||l2!=nil||offset!=0) {
		var temp01,temp02 = 0,0
		if(l1!=nil){temp01=l1.Val}
		if(l2!=nil){temp02=l2.Val}
		currNode03:=ListNode {Val: 0, Next: nil}
		currNode03.Val = (temp01+temp02)%10
		if(count==0){
			l3.Val = currNode03.Val
			l3.Next = nil
			lastNode = &l3
		}else{
			currNode03.Val = (currNode03.Val+offset)%10
			lastNode.Next = &currNode03
			lastNode = lastNode.Next
		}
		if((temp01+temp02+offset)>=10){
			offset=1
		}else{
			offset=0
		}
		count++
		if(l1!=nil){l1 = l1.Next}
		if(l2!=nil){l2 = l2.Next}
	}
	return &l3
}