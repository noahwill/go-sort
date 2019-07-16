package sorts

func selectionSort(slice []int, n int) int {
	compars := 0 /* comparisons */
	var i, j int

	for i = 0; i < n; i++ {
		// Initialize the start of the sort
		min := slice[i]
		minPos := i

		// Min increases if [j] > [i]
		for j = i + 1; j < n; j++ {
			compars++
			if min > slice[j] {
				min = slice[j]
				minPos = j
			}
		}

		// Swap the smallest with the start of the sort
		slice[minPos] = slice[i]
		slice[i] = min
	}

	return compars
}

func insertionSort(slice []int, n int) int {
	compars := 0
	var i, j int
	// Start at [1] because [0:1] is trivially sorted
	// Invariant : the first i elements are sorted
	for i = 1; i < n; i++ {
		compars++
		// Move [i] to its correct position
		// Invariant : the first i+1 elements are sorted
		for j = i; j > 0 && slice[j] < slice[j-1]; j-- {
			temp := slice[j]
			slice[j] = slice[j-1]
			slice[j-1] = temp
			compars++
		}
	}
	return compars
}

func bubbleSort(slice []int, n int) int {
	return 0
}

func quickSort(slice []int, n int) int {
	return 0
}

func shellSort(slice []int, n int) int {
	return 0
}

func mergeSort(slice []int, n int) int {
	return 0
}
