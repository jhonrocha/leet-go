package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	biggest, current := 0, 0
	store := make(map[rune]int)
	for pos, letter := range s {
		if idx, found := store[letter]; found && (pos-idx <= current) {
			current = pos - idx 
		} else {
			current++
			if current > biggest {
				biggest = current
			}
		}
		store[letter] = pos
	}
	return biggest
}

func main() {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("dvdf", lengthOfLongestSubstring("dvdf"))
	fmt.Println("bbtablud", lengthOfLongestSubstring("bbtablud"))
}
