package main

import "fmt"

// main is the entry point of our Go program.
func main() {
	// 1. Declare a variable to store the number we want to check.
	// We use 'int' which helps store integer numbers (whole numbers).
	var numberToCheck int

	// 2. Ask the user for input.
	// fmt.Print prints text to the screen without a new line at the end.
	fmt.Print("Enter a number to check if it's prime: ")

	// 3. Read the number from the user's keyboard.
	// &numberToCheck passes the "address" of the variable so Scan can modify it.
	fmt.Scan(&numberToCheck)

	// Corner case: numbers less than 2 are not prime.
	if numberToCheck < 2 {
		fmt.Printf("%d is NOT a prime number.\n", numberToCheck)
		return
	}

	// 4. Create the Sieve of Eratosthenes.
	// We make a slice (a dynamic array) of booleans (true/false codes).
	// We need indices up to 'numberToCheck', so we make the size 'numberToCheck + 1'.
	// By default, Go sets boolean values to 'false'.
	isPrime := make([]bool, numberToCheck+1)

	// 5. Initialize the sieve.
	// We assume every number from 2 to numberToCheck is prime initially.
	// We set them to 'true'.
	for i := 2; i <= numberToCheck; i++ {
		isPrime[i] = true
	}

	// 6. Apply the Sieve logic.
	// We loop starting from 2.
	// We continue as long as i*i is less than or equal to our number.
	// (Check up to the square root of the number).
	iterationCount := 0
	for i := 2; i*i <= numberToCheck; i++ {
		iterationCount++
		// If isPrime[i] is still true, it means 'i' is a prime number.
		if isPrime[i] {
			// Now we "mark" all multiples of 'i' as NOT prime (false).
			// We start at i*i (since smaller multiples like 2*i would have been handled by 2).
			// We increment by 'i' each time (i*i, i*i+i, i*i+2i, ...).
			for j := i * i; j <= numberToCheck; j += i {
				isPrime[j] = false // Mark as not prime
			}
		}

		if !isPrime[numberToCheck] {
			fmt.Printf("%d is NOT a prime number!\n", numberToCheck)
			fmt.Printf("Number of iterations: %d\n", iterationCount)
			fmt.Printf("The smallest prime factor is: %d\n", i)
			return
		}
	}

	// 7. Output the result.
	// we just check the boolean value at the index of our number.
	if isPrime[numberToCheck] {
		fmt.Printf("%d IS a prime number!\n", numberToCheck)
		fmt.Printf("Number of iterations: %d\n", iterationCount)
	} else {
		fmt.Printf("%d is NOT a prime number.\n", numberToCheck)
		fmt.Printf("Number of iterations: %d\n", iterationCount)
	}
}
