package skiplist

import (
	"math/rand"
	"time"
)

const (
	MAX_L = 32
	P     = 0.5
)

type Node struct {
	v  int      // value
	ls []*Level // node's index
}

type Level struct {
	next *Node
}

type SkipList struct {
	hn *Node // header
	h  int   // height
	c  int   // count
}

func New() *SkipList {
	return &SkipList{
		hn: NewNode(MAX_L, 0),  // head node doesn't count
		h:  1,
		c:  0,
	}
}

func NewNode(level, val int) *Node {
	node := new(Node)
	node.v = val
	node.ls = make([]*Level, level)

	for i := 0; i < len(node.ls); i++ {
		node.ls[i] = new(Level)
	}
	return node
}

/**
get a random level, the odds is
	1/2  = 1
	1/4  = 2
    1/8  = 3
	...
	1/2 + 1/4 + 1/8 ... + 1/n
*/
func (sl *SkipList) randomL() int {
	l := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for r.Float64() < P && l < MAX_L {
		l++
	}
	return l
}

func (sl *SkipList) Add(value int) bool {
	if value <= 0 {
		return false
	}
	update := make([]*Node, MAX_L)

	th := sl.hn
	for i := sl.h - 1; i >= 0; i-- {
		// find the right place in each level
		for th.ls[i].next != nil && th.ls[i].next.v < value {
			th = th.ls[i].next
		}

		// make sure no duplicates
		if th.ls[i].next != nil && th.ls[i].next.v == value {
			return false
		}

		update[i] = th
	}

	// get the level for the new item
	level := sl.randomL()
	node := NewNode(level, value)

	if level > sl.h {
		sl.h = level
	}

	for i := 0; i < level; i++ {
		// concat new node of the upper level to the head node
		if update[i] == nil {
			sl.hn.ls[i].next = node
			continue
		}
		// insert node
		node.ls[i].next = update[i].ls[i].next
		update[i].ls[i].next = node
	}

	sl.c++
	return true
}

func (sl *SkipList) Search(value int) (*Node, bool) {
	var node *Node
	th := sl.hn
	// search for the top level
	for i := sl.h - 1; i >= 0; i-- {
		for th.ls[i].next != nil && th.ls[i].next.v <= value {
			th = th.ls[i].next
		}
		// if no match, then search the next level
		if th.v == value {
			node = th
			break
		}
	}

	if node == nil {
		return nil, false
	}

	return node, true
}

func (sl *SkipList) Delete(value int) bool {
	var node *Node
	last := make([]*Node, sl.h)
	th := sl.hn
	// from top to bottom, delete all match nodes
	for i := sl.h - 1; i >= 0; i-- {
		for th.ls[i].next != nil && th.ls[i].next.v < value {
			th = th.ls[i].next
		}

		last[i] = th
		// find the node to delete
		if th.ls[i].next != nil && th.ls[i].next.v == value {
			node = th.ls[i].next
		}
	}

	// no match
	if node == nil {
		return false
	}

	for i := 0; i < len(node.ls); i++ {
		last[i].ls[i].next = node.ls[i].next
		node.ls[i].next = nil
	}

	// delete empty levels
	for i := 0; i < len(sl.hn.ls); i++ {
		if sl.hn.ls[i].next == nil {
			sl.h = i
			break
		}
	}
	sl.c--
	return true
}

func (sl *SkipList) IsEmpty() bool {
	return sl.c == 0
}

func (sl *SkipList) Count() int {
	return sl.c
}
