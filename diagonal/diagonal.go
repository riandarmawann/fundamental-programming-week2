package main

import (
	"fmt"
	"math"
)

func countDiagonalDifference(arr [][]int) int {
	var diagonal1Sum, diagonal2Sum int

	// Menghitung jumlah diagonal pertama (dari kiri atas ke kanan bawah)
	for i := 0; i < len(arr); i++ {
		diagonal1Sum += arr[i][i]
	}

	// Menghitung jumlah diagonal kedua (dari kanan atas ke kiri bawah)
	for i := 0; i < len(arr); i++ {
		diagonal2Sum += arr[i][len(arr)-1-i]
	}

	// Menghitung selisih antara jumlah diagonal dan menghindari nilai negatif
	difference := int(math.Abs(float64(diagonal1Sum - diagonal2Sum)))

	return difference
}

func main() {
	data := [][]int{
		{1, 2, 5},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(countDiagonalDifference(data))
}
