package main

import (
	"bytes"
	"fmt"
)

func main(){
	var s IntSet
	s.AddAll(9,5,22,55,3)
	s.Add(25)

	fmt.Println(s.String())


	for _, value := range s.Elems() {
		fmt.Println(value)
	}
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

func(s *IntSet) Elems() []uint64 {
	elems := make([]uint64, len(s.words)*64)
	count := 0
	for i, word := range s.words {
		fmt.Println(word)
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems[count] = uint64(64*i+j)
				count+=1
			}
		}
	}
	count = 0
	//strip away zeroes
	for i := len(elems)-1; i >= 0; i-- {
		if(elems[i] == 0){
			count+=1
		}
	}
	return elems[:len(elems)-count]
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string