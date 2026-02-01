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
	numberToCheck, err := strconv.ParseUint(os.Args[1], 10, 64)
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
	// for key := range groupedFactors {
	// 	verifyPrime(key)
	// }

	// Create a formatted string for prime factors
	output := prettifyFactors(groupedFactors)

	// Print factors ungrouped
	// output = arrayJoin(allFactors, " * ")
	// fmt.Printf("Ungrouped factors are:\n %s\n", arrayJoin(allFactors, " * "))

	// Print prettified prime factorial representation
	fmt.Printf("Factors are:\n %s", output)
}

func findTwoFactors(number uint64, start uint64, isPrime []bool) []uint64 {
	// Function to take an integer as an input and return the smallest and biggest factors
	// The smallest factor can only be a prime number (mathemetical certainty)
	// If the number is a prime, it returns 1 and the number itself

	// TODO: Fast integer square root. Potentially useless due to modern hardware sqrt functions
	// limit := uint64(math.Sqrt(float64(number)))
	limit := iSqrt(number)
	factors := make([]uint64, 2)

	for i := start; i <= limit; i++ {
		if isPrime[i] {
			// Checking if i is a factor. Early exit if it is
			if number%i == 0 {
				factors[0] = i
				factors[1] = number / i
				return factors
			}

			// Updating the sieve for future iterations
			for j := i * i; j <= limit; j += i {
				// i * i is where we start with because lower values like 2 * i are already evaluated
				// j <= limit for quick exit. This makes the sieve incomplete, but we do not care about larger values anyway
				isPrime[j] = false
			}
		}
	}

	factors[0] = 1
	factors[1] = number
	return factors
}

func findAllFactors(number uint64) []uint64 {
	// Function that returns all prime factors of a number as an array
	// If input is prime, it returns 1 and the number itself

	var allFactors []uint64

	factors := make([]uint64, 2)
	var small, large uint64
	small = 1
	large = number

	// Biggest prime factor cannot be higher than the sqrt
	// TODO: Fast integer square root. Potentially useless due to modern hardware sqrt functions
	// maxFactor := uint64(math.Sqrt(float64(number)))
	maxFactor := iSqrt(number)

	isPrime := make([]bool, maxFactor+1) // +1 to ensure off by one. 0 is effectively never used

	// Initial condition, assumes all numbers are prime
	var i uint64
	for i = 2; i <= maxFactor; i++ {
		isPrime[i] = true
	}

	// Smallest possible factor
	var start uint64
	start = 2

	for i := 0; i < MAX_ITER; i++ {
		// Reuse same sieve for future computations
		factors = findTwoFactors(large, start, isPrime)
		small = factors[0]
		large = factors[1]

		if large == number {
			factors := []uint64{1, number}
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

func groupPowers(array []uint64) map[uint64]uint64 {
	// Takes a list of numbers and returns a hashmap. The keys are the uniquified numbers and the values are the number of occurrence
	factors := make(map[uint64]uint64)
	var base uint64

	for i := 0; i < len(array); i++ {
		base = array[i]
		factors[base]++
	}

	return factors
}

func prettifyFactors(groupedFactors map[uint64]uint64) string {
	// Takes a hashmap containing prime factors and number of occurences as key-value pairs
	// Returns a prettified stitched together integer factorisation
	var element string
	var output string
	primeFactorCount := len(groupedFactors)
	iterCount := 0

	keys := make([]uint64, 0, len(groupedFactors))

	for k := range groupedFactors {
		keys = append(keys, k)
	}
	// Source - https://stackoverflow.com/a/48568680
	// Posted by vicentazo, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-02-02, License - CC BY-SA 3.0
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

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

func iSqrt(number uint64) uint64 {
	// Initializing with quick lookup table for performance
	if number < 4 {
		return 1
	} else if number < 9 {
		return 2
	} else if number < 16 {
		return 3
	} else if number < 25 {
		return 4
	} else if number < 36 {
		return 5
	} else if number < 49 {
		return 6
	} else if number < 64 {
		return 7
	} else if number < 81 {
		return 8
	} else if number < 100 {
		return 9
	}

	// For 100 and above the square root cannot be > 1/10th of the number
	// The lookup table effectively reduced the subject by an order or magnitude
	// Binary search reduces by 3 iterations (2^3) and a little more
	limit := number / 10
	var low, mid, high uint64
	low = 2
	mid = (limit + 2) / 2
	high = limit

	for (high - low) > 1 {
		if (mid * mid) == number {
			return mid
		} else if (mid * mid) > number {
			high = mid
			mid = (low + high) / 2
		} else if (mid * mid) < number {
			low = mid
			mid = (low + high) / 2
		}
	}

	fmt.Printf("square root of %d is %d\n", number, mid)
	return mid
}

// func verifyPrime(number int) {
// 	bigNumber := new(big.Int)
// 	fmt.Sscan(strconv.Itoa(number), bigNumber)
// 	fmt.Printf("%d is probably prime: %v\n", number, bigNumber.ProbablyPrime(20))
// }
