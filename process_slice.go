// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Process

package main

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// ProcessSlice is a slice of type Process. Use it where you would use []Process.
type ProcessSlice []Process

// SortBy returns a new ordered ProcessSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv ProcessSlice) SortBy(less func(Process, Process) bool) ProcessSlice {
	result := make(ProcessSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortProcessSlice(result, less, 0, n, maxDepth)
	return result
}

// Where returns a new ProcessSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv ProcessSlice) Where(fn func(Process) bool) (result ProcessSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapProcessSlice(rcv ProcessSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapProcessSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapProcessSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownProcessSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapProcessSlice(rcv, first, first+i)
		siftDownProcessSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapProcessSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapProcessSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapProcessSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeProcessSlice(rcv ProcessSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapProcessSlice(rcv, a+i, b+i)
	}
}

func doPivotProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeProcessSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeProcessSlice(rcv, less, m, m-s, m+s)
		medianOfThreeProcessSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeProcessSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapProcessSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapProcessSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapProcessSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeProcessSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeProcessSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortProcessSlice(rcv ProcessSlice, less func(Process, Process) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortProcessSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotProcessSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortProcessSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortProcessSlice(rcv, mhi, b)
		} else {
			quickSortProcessSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortProcessSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortProcessSlice(rcv, less, a, b)
	}
}
