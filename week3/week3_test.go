package week3

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
)

var comparisionCounter = 0
type IntArray []int

func (a IntArray) Len()	int { return len(a) }
func (a IntArray) Less(i, j int) bool { comparisionCounter++; return a[i] < a[j] }
func (a IntArray) Swap(i, j int) { a[j], a[i] = a[i], a[j] }

func ExampleQuickSortFirstPivotSmall() {
	arr := IntArray{1, 6, 4, 13, 2, 5, 8, 7, 3, 12, 11, 10, 9, 14, 16}
	QuickSortFirstPivot(arr)
	fmt.Printf("%d", comparisionCounter)
	// reset comparision Counter after test finishes
	comparisionCounter = 0;
	// Output: 47
}

func ExampleQuickSortFirstPivot() {
	file, err := os.Open("QuickSort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make(IntArray, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	QuickSortFirstPivot(arr)
	fmt.Printf("%d", comparisionCounter)
	// reset comparision Counter after test finishes
	comparisionCounter = 0;
	// Output: 10917609
}


func ExampleQuickSortLastPivot() {
	file, err := os.Open("QuickSort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make(IntArray, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	QuickSortLastPivot(arr)
	fmt.Printf("%d", comparisionCounter)
	// reset comparision Counter after test finishes
	comparisionCounter = 0;
	// Output: 10912940
}

func ExampleQuickSortRandomPivot() {
	file, err := os.Open("QuickSort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make(IntArray, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	QuickSortRandomPivot(arr)
	fmt.Printf("%d", comparisionCounter)
	// reset comparision Counter after test finishes
	comparisionCounter = 0;
	// Output: 10928279
}

func ExampleQuickSortMedianOfThreePivot() {
	file, err := os.Open("QuickSort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make(IntArray, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	QuickSortMedianOfThreePivot(arr)
	fmt.Printf("%d", comparisionCounter)
	// reset comparision Counter after test finishes
	comparisionCounter = 0;
	// Output: 704098
}