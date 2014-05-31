// Package sorter implements a series of sorting algorithms 
// based on the book "Algorithms 4ed" by Sedgewick
// Author: Leesper
// Email: pascal7718@gmail.com 394683518@qq.com

package sorter

import (
	"reflect"
)

// Sortable is an interface for a collection to be sortable
// All sorting algorithms sort collection that is a kind of Sortable
type Sortable interface {
	Length()			int
	Less(i, j int)		bool
	Exchange(i, j int)
}

// multiKeySorter is a kind of Sortable used for sorting according to different keys
// One can write different lesser functions for a user-defined type and call
// By(*lesser*).Sort(*coll*) to sort the collection
type multiKeySorter struct {
	coll	interface{}
	lesser	func(o1, o2 interface{}) bool
}

func (mks *multiKeySorter) Length() int {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return slice.Len()
	}
	panic("passing a non-slice type")
}

func (mks *multiKeySorter) Exchange(i, j int) {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		temp := reflect.ValueOf(slice.Index(i).Interface())
		slice.Index(i).Set(reflect.ValueOf(slice.Index(j).Interface()))
		slice.Index(j).Set(temp)
		return
	}
	panic("passing a non-slice type")
}

func (mks *multiKeySorter) Less(i, j int) bool {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return mks.lesser(slice.Index(i).Interface(), slice.Index(j).Interface())
	}
	panic("passing a non-slice type")
}

// By is a function type for multiple-key sorting
type By func(o1, o2 interface{}) bool

// Sort sorts the slice by lesser func passing in
func (by By) Sort(coll interface{}) {
	mks := &multiKeySorter {
		coll:		coll,
		lesser:		by,
	}
	Shell(mks, 0, mks.Length())
}

// Selection sorts coll[bgn..end)
func Selection(coll Sortable, bgn, end int) {
	for i := bgn; i < end; i++ {
		min := i
		for j := i + 1; j < end && coll.Less(j, min); j++ {
			min = j
		}
		coll.Exchange(i, min)
	}
}

// Insertion sorts coll[bgn..end)
func Insertion(coll Sortable, bgn, end int) {
	for i := bgn; i < end; i++ {
		for j := i; j > bgn && coll.Less(j, j-1); j-- {
			coll.Exchange(j, j - 1)
		}
	}
}

/* The original shell sort proto type
func Shell(coll Sortable) {
	N := coll.Length()
	h := 1
	for h < N / 3 {
		h = 3 * h + 1	// making incremental sequences
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h; j -= h {
				if coll.Less(j, j - h) {
					coll.Exchange(j, j - h)
				}
			}
		}
		h = h / 3
	}
}
*/

// Shell sorts coll[bgn..end)
func Shell(coll Sortable, bgn, end int) {
	h := 1
	for h < (end - bgn) / 3 {
		h = 3 * h + 1	// making incremental sequences
	}
	for h >= 1 {
		for i := h; i < (end - bgn); i++ {
			for j := i; j >= h && coll.Less(j, j - h); j -= h {
				coll.Exchange(j, j - h)
			}
		}
		h = h / 3
	}
}