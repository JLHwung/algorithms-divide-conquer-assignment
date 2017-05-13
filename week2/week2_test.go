package week2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func ExampleCountInversionsInArray_Toy() {
	arr := []int{1, 3, 5, 2, 4, 6}
	fmt.Println(CountInversionsInArray(arr))
	// Output: 3
}

func ExampleCountInversionsInArrayBig() {
	file, err := os.Open("IntegerArray.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	fmt.Println(CountInversionsInArray(arr))
	// Output: 2407905288
}

func BenchmarkCountInversionsInArray(b *testing.B) {
	file, err := os.Open("IntegerArray.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 1024)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, int(number))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CountInversionsInArray(arr)
	}
}
