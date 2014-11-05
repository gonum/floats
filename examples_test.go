// Copyright 2013 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

package floats

import (
	"fmt"
	"math"
)

// Set of examples for all the functions

func ExampleAdd_simple() {
	// Adding three slices together. Note that
	// the result is stored in the first slice
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{5, 6, 7, 8}
	s3 := []float64{1, 1, 1, 1}
	Add(s1, s2)
	Add(s1, s3)

	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
	// Output:
	// s1 = [7 9 11 13]
	// s2 = [5 6 7 8]
	// s3 = [1 1 1 1]
}

func ExampleAdd_newslice() {
	// If one wants to store the result in a
	// new container, just make a new slice
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{5, 6, 7, 8}
	s3 := []float64{1, 1, 1, 1}
	dst := make([]float64, len(s1))

	AddTo(dst, s1, s2)
	Add(dst, s3)

	fmt.Println("dst =", dst)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
	// Output:
	// dst = [7 9 11 13]
	// s1 = [1 2 3 4]
	// s2 = [5 6 7 8]
	// s3 = [1 1 1 1]
}

func ExampleAdd_unequallengths() {
	// If the lengths of the slices are unknown,
	// use Eqlen to check
	s1 := []float64{1, 2, 3}
	s2 := []float64{5, 6, 7, 8}

	eq := EqualLengths(s1, s2)
	if eq {
		Add(s1, s2)
	} else {
		fmt.Println("Unequal lengths")
	}
	// Output:
	// Unequal lengths
}

func ExampleAddConst() {
	s := []float64{1, -2, 3, -4}
	c := 5.0

	AddConst(c, s)

	fmt.Println("s =", s)
	// Output:
	// s = [6 3 8 1]
}

func ExampleCumProd() {
	s := []float64{1, -2, 3, -4}
	dst := make([]float64, len(s))

	CumProd(dst, s)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [1 -2 -6 24]
	// s = [1 -2 3 -4]
}

func ExampleCumSum() {
	s := []float64{1, -2, 3, -4}
	dst := make([]float64, len(s))

	CumSum(dst, s)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [1 -1 2 -2]
	// s = [1 -2 3 -4]
}

func ExampleFilter_simple() {
	s := []float64{1, 3, 2, 5, 6}
	f := func(a float64) bool {
		return a > 3
	}

	dst, _ := Filter(f, s, -1)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [5 6]
	// s = [1 3 2 5 6]
}

func ExampleFilter_susequence() {
	// Let's use state magic to filter elements
	// so the resulting list is a strictly increasing
	// subsequence!

	max := math.Inf(-1)
	s := []float64{1, 3, 2, 6, 5}
	f := func(a float64) bool {
		if a > max {
			max = a
			return true
		}

		return false
	}

	dst, _ := Filter(f, s, -1)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [1 3 6]
	// s = [1 3 2 6 5]
}

func ExampleFoldRight() {
	s := []float64{9, -2, 21, 4}
	f := func(a, b float64) float64 {
		if a < b {
			return a * b
		}

		return a - b
	}
	initial := 5.0

	val := FoldRight(f, s, initial)

	fmt.Println("val =", val)
	fmt.Println("initial =", initial)
	fmt.Println("s =", s)
	// Output:
	// val = 11
	// initial = 5
	// s = [9 -2 21 4]
}

func ExampleFoldLeft() {
	s := []float64{9, -2, 21, 4}
	f := func(a, b float64) float64 {
		if a < b {
			return a * b
		}

		return a - b
	}
	initial := 5.0

	val := FoldLeft(f, s, initial)

	fmt.Println("val =", val)
	fmt.Println("initial =", initial)
	fmt.Println("s =", s)
	// Output:
	// val = 22
	// initial = 5
	// s = [9 -2 21 4]
}
