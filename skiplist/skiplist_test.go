package skiplist

import (
	"math/rand"
	"testing"
)

const (
	maxN = 10
)

func TestAddAndSearch(t *testing.T) {
	var list = New()

	// Test at the beginning of the list.
	for i := 1; i < maxN; i++ {
		list.Add(maxN - i)
	}
	for i := 1; i < maxN; i++ {
		if _, ok := list.Search(maxN - i); !ok {
			t.Fail()
		}
	}

	list = New()
	// Test at the end of the list.
	for i := 1; i < maxN; i++ {
		list.Add(i)
	}
	for i := 1; i < maxN; i++ {
		if _, ok := list.Search(i); !ok {
			t.Fail()
		}
	}

	list = New()
	// Test at random positions in the list.
	rList := rand.Perm(maxN)
	for _, e := range rList {
		list.Add(e)
	}
	for _, e := range rList {
		if _, ok := list.Search(e); !ok {
			t.Fail()
		}
	}

}

func TestDelete(t *testing.T) {

	var list = New()

	// Delete elements at the beginning of the list.
	for i := 1; i < maxN; i++ {
		list.Add(i)
	}
	for i := 1; i < maxN; i++ {
		list.Delete(i)
	}
	if !list.IsEmpty() {
		t.Fail()
	}

	list = New()
	// Delete elements at the end of the list.
	for i := 1; i < maxN; i++ {
		list.Add(i)
	}
	for i := 1; i < maxN; i++ {
		list.Delete(maxN - i - 1)
	}
	if list.Count() != 1 {
		t.Fail()
	}

	list = New()
	// Delete elements at random positions in the list.
	rList := rand.Perm(maxN)
	for _, e := range rList {
		list.Add(e)
	}
	for _, e := range rList {
		list.Delete(e)
	}
	if !list.IsEmpty() {
		t.Fail()
	}
}

