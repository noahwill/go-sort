package sorts

func (s slice) swap(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s slice) SelectionSort(n int) int {
	compars := 0 /* comparisons */
	// Select the first element as the smallest known
	// element
	// Invariant : the first i elements are sorted
	for i := 0; i < n; i++ {
		min := s[i]
		minPos := i
		// Min increases if [j] > [i]
		// Invariant : min is the smallest value in [i:j]
		for j := i + 1; j < n; j++ {
			compars++
			if min > s[j] {
				min = s[j]
				minPos = j
			}
		}
		// Swap the smallest with the start of the sort
		s[minPos] = s[i]
		s[i] = min
	}
	return compars
}

func (s slice) InsertionSort(n int) int {
	compars := 0
	// Start at [1] because [0:1] is trivially sorted
	// Invariant : the first i elements are sorted
	for i := 1; i < n; i++ {
		compars++
		// Move [i] to its correct position
		// Invariant : the first i+1 elements are sorted
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s.swap(j, j-1)
			compars++
		}
	}
	return compars
}

func (s slice) BubbleSort(n int) int {
	compars := 0
	changed := true
	for changed {
		changed = false
		// Invariant : after i iterations, the right most i elements
		// are sorted and in place
		for i := 1; i < n; i++ {
			compars++
			if s[i] < s[i-1] {
				changed = true
				s.swap(i, i-1)
			}
		}
	}
	return compars
}

// Precondition  : 0 <= start && start <= stop && stop <= len(s)
// Postcondition : s[start:stop] is sorted
func (s slice) quickSortR(start int, stop int) int {
	compars := 0
	if stop-start > 1 {
		pivot := s[stop-1]
		i := start - 1
		for j := start; j < stop-1; j++ {
			if s[j] <= pivot {
				s.swap(j, i+1)
				i++
				compars++
			}
		}
		s.swap(stop-1, i+1)
		s.quickSortR(start, i+1)
		s.quickSortR(i+2, stop)
		return compars
	}
	return compars
}

func (s slice) QuickSort(n int) int {
	return s.quickSortR(0, n)
}

func (s slice) ShellSort(n int) int {
	compars := 0
	for gap := n / 3; ; gap = gap/13 + 1 {
		for k := 0; k < gap; k++ {
			for i := k + gap; i < n; i += gap {
				compars++
				for j := i; j >= gap && s[j] < s[j-gap]; j -= gap {
					s.swap(j, j-gap)
					compars++
				}
			}
			if gap == 1 {
				break
			}
		}
	}
	return compars
}

// Precondition  : s[0:start-1] and s[start:stop-1] are sorted
// Postcondition : the first start + stop elements have been sorted
func (s slice) merge(h slice, start, stop int) int {
	n := stop - start
	midpoint := (start + stop) / 2
	compars := 0
	// sub slices
	i := start               // index to the first sub slice
	j := midpoint            // index to the second sub slice
	for k := 0; k < n; k++ { // index to the helper slice
		if i >= midpoint {
			h[k] = s[j]
			j++
		} else if j >= stop {
			h[k] = s[i]
			i++
		} else if s[i] < s[j] {
			compars++
			h[k] = s[i]
			i++
		} else {
			compars++
			h[k] = s[j]
			j++
		}
	}
	return compars
}

func (s slice) mergeSortR(h slice, start, stop int) int {
	if start+1 >= stop {
		return 0
	}

	compars := 0
	midpoint := (start + stop) / 2
	n := stop - start

	compars += s.mergeSortR(h, start, midpoint)
	compars += s.mergeSortR(h, midpoint, stop)
	compars += s.merge(h, start, stop)

	for k := 0; k < n; k++ {
		s[start+k] = h[k]
	}

	return compars
}

func (s slice) MergeSort(n int) int {
	h := make(slice, len(s), len(s))
	compars := s.mergeSortR(h, 0, n)
	return compars
}
