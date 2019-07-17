package sorts

import (
	'fmt'
	'math/rand'
	'time'
)

var initialized = false

type slice []int

func RandomSlice(s slice, n int) {
	if !initialized {
		rand.Seed(time.Now().UTC().UnixNano())
		initialized = true
	}
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(999)
	}
}

func BlankSlice(s slice, n int) {
	for i := 0; i < n; i++ {
		s[i] = 0
	}
}

func DisplaySlice(s slice, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf('%d', s[i])
		fmt.Printf('\n')
	}
}

func GetTime() time.Time {
	return time.Now()
}

func CopySlice(s slice, h slice, n int) {
	for i := 0; i < n; i++ {
		h[i] = s[i]
	}
}

func IsSorted(s slice, n int) bool {
	sorted := true
	for i := 0; i < n-1; i++ {
		sorted = (s[i] <= s[i+1])
	}
	return sorted
}

func SameContents(s slice, h slice, n int) bool {
	return false
}

