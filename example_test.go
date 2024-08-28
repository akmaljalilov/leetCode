package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	res := twoSum(nums, 6)
	if len(res) != 2 || nums[res[0]]+nums[res[1]] != 6 {
		t.Error("Invalid response")
	}
}

func TestIsPalindrome(t *testing.T) {
	if isPalindrome(221) {
		t.Error("Invalid response")
	}
	if !isPalindrome(9999) {
		t.Error("Invalid response")
	}
	if !isPalindrome(8) {
		t.Error("Invalid response")
	}
	if !isPalindrome(121) {
		t.Error("Invalid response")
	}
	if isPalindrome(500) {
		t.Error("Invalid response")
	}
	if isPalindrome(-121) {
		t.Error("Invalid response")
	}
}
func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Error("Invalid response")
	}
	if romanToInt("IV") != 4 {
		t.Error("Invalid response")
	}
	if romanToInt("LVIII") != 58 {
		t.Error("Invalid response")
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Error("Invalid response")
	}
}
func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		t.Error("Invalid response")
	}
	if longestCommonPrefix([]string{"dog", "racecar", "car"}) != "" {
		t.Error("Invalid response")
	}
}

func TestIsValid(t *testing.T) {
	if isValid("{") {
		t.Error("Invalid response")
	}
	if !isValid("{[]}") {
		t.Error("Invalid response")
	}
	//if isValid("([)]") {
	//	t.Error("Invalid response")
	//}
	//if !isValid("()[]{}") {
	//	t.Error("Invalid response")
	//}
	//if !isValid("()") {
	//	t.Error("Invalid response")
	//}
	//if isValid("(]") {
	//	t.Error("Invalid response")
	//}
}

func TestMergeTwoLists(t *testing.T) {
	t.Run("", func(t *testing.T) {
		list1 := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			}}
		list2 := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			}}
		mergeTwoLists(&list1, &list2)
	})
	t.Run("", func(t *testing.T) {
		list1 := ListNode{
			Val: 5,
		}
		list2 := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			}}
		mergeTwoLists(&list1, &list2)
	})
	t.Run("", func(t *testing.T) {
		mergeTwoLists(nil, nil)
	})

}
