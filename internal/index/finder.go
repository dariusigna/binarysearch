package index

import (
	"math"
)

// Finder is a type that finds the index of a value in a sorted array.
type Finder struct {
	values            []int
	conformationLevel int
}

// NewFinder creates a new Finder with the given values.
func NewFinder(values []int, conformationLevel int) *Finder {
	return &Finder{values: values, conformationLevel: conformationLevel}
}

// FindIndex finds the index of the value in the sorted array.
func (f *Finder) FindIndex(value int) (int, bool) {
	if len(f.values) == 0 {
		return -1, false
	}

	left, right := 0, len(f.values)-1
	for left <= right {
		mid := left + (right-left)/2 // (left + right) / 2 might overflow
		if f.values[mid] == value {
			return mid, true
		}

		if value < f.values[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// left = right + 1
	// Exact match not found -> check closest index for conformation percentage
	closestIndex := f.getClosestIndex(left, right, value)
	if !checkConformation(value, f.values[closestIndex], f.conformationLevel) {
		return -1, false

	}
	return closestIndex, true
}

func checkConformation(value, comparedValue, percent int) bool {
	diff := math.Abs(float64(value - comparedValue))
	return diff <= float64(value)/float64(percent)
}

func (f *Finder) getClosestIndex(left, right, value int) int {
	closestIndex := -1
	// Check if on bounds
	if left == len(f.values) {
		closestIndex = right
	} else if right == -1 {
		closestIndex = left
	}

	if closestIndex != -1 {
		return closestIndex
	}

	// Check if the value is closer to the left or right
	if f.values[left]-value < value-f.values[right] {
		closestIndex = left
	} else {
		closestIndex = right
	}

	return closestIndex
}
