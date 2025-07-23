// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, दुनिया!")
	fmt.Println(Fizzbuzz(15))
	fmt.Println(IsPrime(-1))
	fmt.Println(IsPalindrome("ab"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	if n % 3  == 0 && (n % 5 == 0) {
	   return "FizzBuzz"
	}
	if n % 3 == 0 {
	   return "Fizz"
	}
	if n % 5 == 0 {
	   return "Buzz"
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n < 1 {
		return false
	} else if n == 1{
		return true
	}
	for i:= 2 ; i < n ; i ++ {
		if n % i == 0{
			return true
		}
	} 
	return false
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
    start:=0 
	end:= len(s) - 1
	for i:= 0 ; i < len(s) / 2 ; i ++{
		if s[start + i] != s[end - i]{
			return false
		}
		start += 1
		end -= 1
	}
	return true
}