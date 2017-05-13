package week2

// countSplitInversionsInArray will count the split inversion in sorted left and right, will save sorted array on arr
func countSplitInversionsInArray(left, right, arr []int) int {
	splitInversions := 0
	for i, j, arrIndex := 0, 0, 0; arrIndex < len(arr); arrIndex++ {
		if i == len(left) {
			goto useRight
		}
		if j == len(right) {
			goto useLeft
		}
		if left[i] > right[j] {
			goto useRight
		} else {
			goto useLeft
		}
	useLeft:
		arr[arrIndex] = left[i]
		i++
		continue
	useRight:
		splitInversions += len(left) - i
		arr[arrIndex] = right[j]
		j++
		continue
	}
	return splitInversions
}

// CountInversionsInArray will count the inversions in []int array arr. An inversions is any pair (i, j) such that
// arr[i] > arr[j]
func CountInversionsInArray(arr []int) int {
	length := len(arr)
	// early return when only one/zero element exists
	if length <= 1 {
		return 0
	}
	if length == 2 {
		if arr[0] > arr[1] {
			// sort arr for split inversion counting
			arr[0], arr[1] = arr[1], arr[0]
			return 1
		}
		return 0
	}

	// split array into left and right
	arrCopy := make([]int, length, length)
	copy(arrCopy, arr)
	left, right := arrCopy[:length/2], arrCopy[length/2:]

	leftInversions := CountInversionsInArray(left)
	rightInversions := CountInversionsInArray(right)
	// now left/right should be sorted
	splitInversions := countSplitInversionsInArray(left, right, arr)

	return splitInversions + leftInversions + rightInversions
}
