package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrimeNumber(arr []int) []int {
	var primeNumbers []int

	for _, num := range arr {
		if isPrime(num) {
			primeNumbers = append(primeNumbers, num)
		}
	}

	return primeNumbers
}

func main() {
	fmt.Println(getPrimeNumber([]int{1, 5, 8, 9, 2, 3, 5, 9, 12, 11, 21}))
}
