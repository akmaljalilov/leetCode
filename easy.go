package main

import (
	"math"
	"strings"
)

// https://leetcode.com/problemset/?sorting=W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZXJCeSI6IkRJRkZJQ1VMVFkifV0%3D

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
func twoSum(nums []int, target int) []int {
	newNums := map[int]int{}
	for i := 0; i < len(nums); i++ {
		newNums[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		expectedNumber := target - nums[i]
		if j, ok := newNums[expectedNumber]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}

// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.
//
// Example 1:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	oldX := x
	newX := 0
	oldLength := int(math.Pow10(int(math.Log10(float64(x)))))
	length := oldLength
	for x > 0 {
		i := x / length
		newX += i * oldLength / length
		x = x % length
		length /= 10
	}
	return newX == oldX
}

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
// Example 1:
//
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
// Example 2:
//
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
// Example 3:
//
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func romanToInt(s string) int {
	mapsItem := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sItems := strings.Split(s, "")
	sInt := -1
	newSInt := 0
	sNewItem := ""
	prevKey := -1
	for i := 0; i < len(sItems); i++ {
		if sItem, ok := mapsItem[sItems[i]]; ok {
			if prevKey != -1 && prevKey > sItem {
				newSInt += sInt
				sNewItem = ""
				sInt = -1
			}
			prevKey = sItem
			sNewItem += sItems[i]
			if sInt < 0 {
				sInt = sItem
				continue
			}
			if sInt < sItem {
				sItem -= sInt
				sInt = sItem
			} else {
				sInt += sItem
			}
		}
	}

	return newSInt + sInt
}

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
func longestCommonPrefix(strs []string) string {
	prefix := ""
	length := len(strs)
	if length == 0 {
		return prefix
	}
	strItem := strings.Split(strs[0], "")
	revPrefix := ""
	for i := 0; i < len(strItem); i++ {
		prefix += strItem[i]
		countPrefix := 1
		for j := 1; j < length; j++ {
			if strings.HasPrefix(strs[j], prefix) {
				countPrefix++
			} else {
				break
			}
		}
		if countPrefix == length {
			revPrefix = prefix
		} else {
			break
		}
	}

	return revPrefix
}

// 20. Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
func isValid(s string) bool {
	arr := []rune(s)
	parentheses := make([]int32, 0)
	parenthesesMap := map[int32]int32{
		40:  41,
		91:  93,
		123: 125,
	}
	valid := true
	for _, ch := range arr {
		switch ch {
		case 40, 91, 123:
			parentheses = append(parentheses, ch)
		case 41, 93, 125:
			lenPers := len(parentheses)
			valid = lenPers != 0
			if valid {
				lastP := parentheses[lenPers-1]
				parentheses = parentheses[0 : lenPers-1]
				valid = parenthesesMap[lastP] == ch
			}
		}
		if !valid {
			break
		}
	}
	return valid && len(parentheses) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
21. Merge Two Sorted Lists
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1 2 4
	// 1 3 4
	// 1 1 2 4 4
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var listNewNext1 *ListNode
	var listNew *ListNode
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 == nil {
			if listNewNext1 == nil {
				listNewNext1 = &ListNode{Val: list1.Val, Next: list1.Next}
			} else {
				listNewNext1.Val = list1.Val
				listNewNext1.Next = list1.Next
				listNewNext1 = listNewNext1.Next
			}
		} else if list1 == nil && list2 != nil {
			if listNewNext1 == nil {
				listNewNext1 = &ListNode{Val: list2.Val, Next: list2.Next}
			} else {
				listNewNext1.Val = list2.Val
				listNewNext1.Next = list2.Next
				listNewNext1 = listNewNext1.Next
			}
		} else if list1.Val <= list2.Val {
			listNewNext := *list1.Next
			list1.Next = &ListNode{
				Val:  list2.Val,
				Next: &listNewNext,
			}
			if listNew == nil {
				listNew = &ListNode{Val: list1.Val, Next: list1.Next}
				listNewNext1 = listNew.Next
			} else {
				listNewNext1.Val = list1.Val
				listNewNext1.Next = list1.Next
				listNewNext1 = listNewNext1.Next
			}
			list1 = list1.Next
			if list2 != nil {
				list2 = list2.Next
			}
		} else {
			if listNew == nil {
				listNew = &ListNode{Val: list2.Val, Next: list2.Next}
				listNewNext1 = listNew.Next
			} else {
				listNewNext1.Val = list2.Val
				listNewNext1.Next = list2.Next
				listNewNext1 = listNewNext1.Next
			}
			list2 = list2.Next
		}
	}

	return listNew
}
