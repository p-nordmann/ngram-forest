package pforest

import (
	"testing"
)

func getNGrams(node Node) map[string]int {
	ngrams := make(map[string]int)
	for token, child := range node.children {
		ngrams[string(token)] = child.count
		childNGrams := getNGrams(*child)
		for ngram, count := range childNGrams {
			ngrams[string(token)+ngram] = count
		}
	}
	return ngrams
}

func compare(ngrams1, ngrams2 map[string]int) bool {
	if len(ngrams1) != len(ngrams2) {
		return false
	}
	for ngram, count1 := range ngrams1 {
		count2, ok := ngrams2[ngram]
		if !ok {
			return false
		}
		if count1 != count2 {
			return false
		}
	}
	return true
}

func TestLearn(t *testing.T) {
	text := "Hello world! :)"
	forest := New(2)
	forest.Learn(text)
	ngrams := getNGrams(forest.parent)
	if !compare(ngrams, map[string]int{
		"H":  1,
		"e":  1,
		"l":  3,
		"o":  2,
		" ":  2,
		"w":  1,
		"r":  1,
		"d":  1,
		"!":  1,
		":":  1,
		")":  1,
		"He": 1,
		"el": 1,
		"ll": 1,
		"lo": 1,
		"o ": 1,
		" w": 1,
		"wo": 1,
		"or": 1,
		"rl": 1,
		"ld": 1,
		"d!": 1,
		"! ": 1,
		" :": 1,
		":)": 1,
	}) {
		t.Errorf("Incorrect ngram counts: %v", ngrams)
	}
}

func TestSum(t *testing.T) {
	text1 := "Hello worlllld! :)"
	text2 := "This is a test."

	f1 := New(2)
	f1.Learn(text1)

	f2 := New(3)
	f2.Learn(text2)

	f3 := Sum(f1, f2)
	f4 := New(2)
	f4.Learn(text1)
	f4.maxDepth = 3
	f4.Learn(text2)

	ngrams3 := getNGrams(f3.parent)
	ngrams4 := getNGrams(f4.parent)
	if !compare(ngrams3, ngrams4) {
		t.Error("Unexpected sum results.")
	}
}

func TestIntersection(t *testing.T) {
	text1 := "Hello world! :)"
	text2 := "This is a test."

	f1 := New(2)
	f1.Learn(text1)

	f2 := New(3)
	f2.Learn(text2)

	ngrams := getNGrams(Intersection(f1, f2).parent)
	if !compare(ngrams, map[string]int{
		"e": 1,
		" ": 2,
	}) {
		t.Errorf("Incorrect ngram counts: %v", ngrams)
	}
}

func TestUnionIntersectionDistributivity(t *testing.T) {
	text1 := "Hello world! :)"
	text2 := "This is a test."
	text3 := "My tailor is rich!"

	f1 := New(2)
	f1.Learn(text1)
	f2 := New(3)
	f2.Learn(text2)
	f3 := New(4)
	f3.Learn(text3)

	f5 := Intersection(Union(f1, f2), f3)
	f6 := Union(Intersection(f1, f3), Intersection(f2, f3))

	if ngrams5, ngrams6 := getNGrams(f5.parent), getNGrams(f6.parent); !compare(ngrams5, ngrams6) {
		t.Errorf("Product should be distributive over Sum. %v != %v", ngrams5, ngrams6)
	}

	f7 := Union(Intersection(f1, f2), f3)
	f8 := Intersection(Union(f1, f3), Union(f2, f3))

	if !compare(getNGrams(f7.parent), getNGrams(f8.parent)) {
		t.Error("Sum should be distributive over Product.")
	}
}
