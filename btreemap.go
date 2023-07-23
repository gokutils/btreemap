package btreemap

import (
	"github.com/google/btree"
)

type Key[K any] interface {
	Less(K) bool
}

type item[K Key[K], V any] struct {
	Key   K
	Value V
}

func (left *item[K, V]) Less(right *item[K, V]) bool {
	return left.Key.Less(right.Key)
}

type Map[K Key[K], V any] struct {
	tree *btree.BTreeG[*item[K, V]]
}

func less[K Key[K]](left K, right K) bool {
	return left.Less(right)
}

/*
Mew creates a new B-Tree with the given degree. For example, New(2) will create a 2-3-4 tree (each node contains 1-3 items and 2-4 children).
*/
func New[K Key[K], V any](degree int) *Map[K, V] {
	return &Map[K, V]{
		tree: btree.NewG[*item[K, V]](degree, less[*item[K, V]]),
	}
}

/*
Get looks for the key item in the tree, returning it. It returns (zeroValue, false) if unable to find that item.
*/
func (impl *Map[K, V]) Get(key K) (V, bool) {
	tmp := &item[K, V]{
		Key: key,
	}
	if v, ok := impl.tree.Get(tmp); ok {
		return v.Value, true
	}
	var noop V
	return noop, false
}

/*
Set adds the given item to the tree. If an item in the tree already equals the given one, it is removed from the tree and returned, and the second return value is true. Otherwise, (zeroValue, false)
*/
func (impl *Map[K, V]) Set(key K, value V) (V, bool) {
	tmp := &item[K, V]{
		Key:   key,
		Value: value,
	}
	if v, ok := impl.tree.ReplaceOrInsert(tmp); ok {
		return v.Value, true
	}
	var noop V
	return noop, false
}

/*
Delete removes an item equal to the passed in item from the tree, returning it. If no such item exists, returns (zeroValue, false)
*/
func (impl *Map[K, V]) Delete(key K) (V, bool) {
	tmp := &item[K, V]{
		Key: key,
	}
	if v, ok := impl.tree.Delete(tmp); ok {
		return v.Value, true
	}
	var noop V
	return noop, false
}

/*
Clear removes all items from the btree
*/
func (impl *Map[K, V]) Clear() {
	impl.tree.Clear(false)
}

/*
Len returns the number of items currently in the tree
*/
func (impl *Map[K, V]) Len() int {
	return impl.tree.Len()
}

/*
ForEach calls the iterator for every value in the tree within the range [first, last], until iterator returns false
*/
func (impl *Map[K, V]) ForEach(iterator func(key K, value V) bool) {
	impl.tree.Ascend(func(item *item[K, V]) bool {
		return iterator(item.Key, item.Value)
	})
}
