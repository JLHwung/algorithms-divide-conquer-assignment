package week3

import (
	"math/rand"
	"sort"
)

// quickSort does pivot and recurse into partitioned sort.Interface
func quickSort(data sort.Interface, lo, hi int, pivotChoosing func(data sort.Interface, lo, hi int) int) {
	if hi-lo > 1 {
		pivot := doPivot(data, lo, hi, pivotChoosing)
		quickSort(data, lo, pivot-1, pivotChoosing)
		quickSort(data, pivot+1, hi, pivotChoosing)
	}
}

// doPivot sort around the pivot and returned the choosing pivot
func doPivot(data sort.Interface, lo, hi int, pivotChoosing func(data sort.Interface, lo, hi int) int) int {
	pivot := pivotChoosing(data, lo, hi)
	// move pivot to lo
	if pivot != lo {
		data.Swap(pivot, lo)
	}
	i := lo + 1
	for j := i; j < hi; j++ {
		if data.Less(j, pivot) && i < hi {
			data.Swap(j, i)
			i++
		}
	}
	data.Swap(lo, i-1)
	return i - 1
}

// QuickSort sort data using given pivotChoosing function
func QuickSort(data sort.Interface, pivotChoosing func(data sort.Interface, lo, hi int) int) {
	quickSort(data, 0, data.Len(), pivotChoosing)
}

// QuickSortFirstPivot sort data using first element as pivot
func QuickSortFirstPivot(data sort.Interface) {
	QuickSort(data, func(data sort.Interface, lo, hi int) int {
		return lo
	})
}

// QuickSortLastPivot sort data using last element as pivot
func QuickSortLastPivot(data sort.Interface) {
	QuickSort(data, func(data sort.Interface, lo, hi int) int {
		return hi - 1
	})
}

// QuickSortRandomPivot sort data using random element as pivot
func QuickSortRandomPivot(data sort.Interface) {
	QuickSort(data, func(data sort.Interface, lo, hi int) int {
		return rand.Intn(hi - lo) + lo
	})
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data sort.Interface, m1, m0, m2 int) {
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] <= data[m2]
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
}

// QuickSortMedianOfThreePivot sort data using median of three as pivot
func QuickSortMedianOfThreePivot(data sort.Interface) {
	QuickSort(data, func(data sort.Interface, lo, hi int) int {
		m := lo + (hi - 1 - lo) / 2
		medianOfThree(data, lo, hi - 1, m)
		return m
	})
}
