// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	str := "wwweee"
	str2 := "eeewww"
	item := Anagram(str, str2)
	fmt.Println(item)
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//
//	ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func isnumber(number byte) bool {
	if int(number) >= 48 && int(number) <= 57 {
		return true
	}
	return false
}
func ParsePhone(phone string) string {
	// TODO
	var count int = 0
	var new_phone string = "("
	for i := 0; i < len(phone); i++ {
		if isnumber(phone[i]) {
			new_phone += string(phone[i])
			count++
			if count == 3 {
				new_phone += ") "
			}
			if count == 6 {
				new_phone += "-"
			}
		}

	}
	return new_phone
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	// TOD
	smap := make(map[rune]int)
	var ok bool = true
	for _, value := range s1 {
		smap[value] += 1
	}
	for _, value := range s2 {
		if _, ok = smap[value]; !ok {
			return false
		}
		if smap[value] == 0 {
			return false
		}
		smap[value] -= 1
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	// TODO
	var arr []int
	for i := 0; i < len(e); i++ {
		if e[i]%2 == 0 {
			arr = append(arr, e[i])
		}
	}
	return arr
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	// TODO
	sum := 1
	for i := 0; i < len(e); i++ {
		sum *= e[i]
	}
	return sum
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	// TODO
	maps := make(map[int]int)
	var unique []int
	for i := 0; i < len(e); i++ {
		maps[e[i]] += 1
	}
	for key, value := range maps {
		if value == 1 {
			unique = append(unique, key)
		}
	}
	return unique
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	// TODO
	maps := make(map[int]string)
	for key, value := range kv {
		maps[value] = key
	}
	return maps
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	// TODO
	maps := make(map[rune]int)
	for _, r := range s {
		maps[r] += 1
	}
	new_maps := make(map[rune]int)
	for key, value := range maps {
		if value > k {
			new_maps[key] = value
		}
	}
	return new_maps
}
