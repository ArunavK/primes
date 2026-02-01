package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MAX_ITER int = 1000

func main() {
	numberToCheck, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Argument is not a valid integer! %v", err)
	}

	// Check user input
	// fmt.Println("Enter a number to find its factors: ")
	// var numberToCheck int
	// fmt.Scan(&numberToCheck)

	// Handle 0 and 1
	if numberToCheck < 2 {
		fmt.Printf("%d does not have any factors.\n", numberToCheck)
		return
	}

	// Find all factors in a list
	allFactors := findAllFactors(numberToCheck)

	// Handle prime number inputs
	if allFactors[0] == 1 {
		fmt.Printf("%d is a prime number!\n", numberToCheck)
		return
	}

	// Group all prime factors into a hashmap. The bases are keys and exponents are values
	groupedFactors := groupPowers(allFactors)

	// Create a formatted string for prime factors
	output := prettifyFactors(groupedFactors)

	// Print factors ungrouped
	// output = arrayJoin(allFactors, " * ")
	// fmt.Printf("Ungrouped factors are:\n %s\n", arrayJoin(allFactors, " * "))

	// Print prettified prime factorial representation
	fmt.Printf("Factors are:\n %s", output)
}

func findTwoFactors(number int, start int, isPrime []bool) []int {
	// Function to take an integer as an input and return the smallest and biggest factors
	// The smallest factor can only be a prime number (mathemetical certainty)
	// If the number is a prime, it returns 1 and the number itself
	factors := make([]int, 2)
	// isPrime := make([]bool, number+1)

	// for i := 2; i <= number; i++ {
	// 	isPrime[i] = true
	// }

	for i := start; i*i <= number; i++ {
		if isPrime[i] {
			for j := i * i; j <= number; j += i {
				isPrime[j] = false
			}
		}

		if !isPrime[number] {
			factors[0] = i
			factors[1] = number / i
			return factors
		}
	}

	factors[0] = 1
	factors[1] = number
	return factors
}

func findAllFactors(number int) []int {
	// Function that returns all prime factors of a number as an array
	// If input is prime, it returns 1 and the number itself

	if number < 2 {
		fmt.Printf("%d does not have any factors.\n", number)
		factors := []int{1, number}
		return factors
	}

	var allFactors []int
	factors := make([]int, 2)

	small := 1
	large := number

	isPrime := make([]bool, number+1)
	for i := 2; i <= number; i++ {
		isPrime[i] = true
	}
	start := 2

	for i := 0; i < MAX_ITER; i++ {
		// Reuse same sieve for future computations
		factors = findTwoFactors(large, start, isPrime)
		small = factors[0]
		large = factors[1]

		if large == number {
			factors := []int{1, number}
			return factors
		}

		if small == 1 {
			allFactors = append(allFactors, large)
			break
		}
		allFactors = append(allFactors, small)
		start = small
	}
	return allFactors
}

func arrayJoin(array []int, delim string) string {
	// Function to join an integer array with a delimiter and return it as a string
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), delim), "[]")
}

func groupPowers(array []int) map[int]int {
	// Takes a list of numbers and returns a hashmap. The keys are the uniquified numbers and the values are the number of occurrence
	factors := make(map[int]int)
	var base int

	for i := 0; i < len(array); i++ {
		base = array[i]
		factors[base]++
	}

	return factors
}

func prettifyFactors(groupedFactors map[int]int) string {
	// Takes a hashmap containing prime factors and number of occurences as key-value pairs
	// Returns a prettified stitched together integer factorisation
	var element string
	var output string
	primeFactorCount := len(groupedFactors)
	iterCount := 0

	keys := make([]int, 0, len(groupedFactors))

	for k := range groupedFactors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		base := k
		exponent := groupedFactors[k]
		iterCount++
		if exponent == 1 {
			element = fmt.Sprintf("%d", base)
		} else {
			element = fmt.Sprintf("%d^%d", base, exponent)
		}
		if iterCount != primeFactorCount {
			element += " * "
		}
		output += element
	}
	return output
}
