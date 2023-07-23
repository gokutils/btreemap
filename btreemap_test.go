package btreemap_test

import (
	"testing"

	"github.com/gokutils/btreemap"
	"github.com/stretchr/testify/assert"
)

type IntLess int

func (v IntLess) Less(right IntLess) bool {
	return v < right
}

func TestNew(t *testing.T) {
	tree := btreemap.New[IntLess, string](10)
	v, ok := tree.Set(IntLess(1), "1")
	assert.False(t, ok)
	assert.Equal(t, v, "")
	v, ok = tree.Set(IntLess(123), "123")
	assert.False(t, ok)
	assert.Equal(t, v, "")
	v, ok = tree.Set(IntLess(10), "10")
	assert.False(t, ok)
	assert.Equal(t, v, "")
	v, ok = tree.Set(IntLess(10), "100")
	assert.True(t, ok)
	assert.Equal(t, v, "10")
	// get
	v, ok = tree.Get(IntLess(10))
	assert.True(t, ok)
	assert.Equal(t, v, "100")
	v, ok = tree.Get(IntLess(20))
	assert.False(t, ok)
	assert.Equal(t, v, "")
	// delete
	v, ok = tree.Delete(IntLess(123))
	assert.True(t, ok)
	assert.Equal(t, v, "123")
	v, ok = tree.Delete(IntLess(123))
	assert.False(t, ok)
	assert.Equal(t, v, "")
	//forEach
	count := 0
	tree.ForEach(func(key IntLess, value string) bool {
		count += 1
		return true
	})
	assert.Equal(t, count, tree.Len())
	tree.Clear()
	assert.Equal(t, 0, tree.Len())
}
