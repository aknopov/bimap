package bimap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringify[V comparable](value V, ok bool) string {
	if ok {
		return fmt.Sprintf("%v", value)
	}
	return "nothing!"
}

func TestBiMapBasics(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()
	assert.Equal(0, aBimap.Size())

	aBimap.Put("Hello", 1)
	assert.Equal(1, aBimap.Size())

	assert.True(aBimap.ContainsKey("Hello"))
	assert.True(aBimap.ContainsValue(1))
	assert.False(aBimap.ContainsKey("guy"))

	fmt.Println(stringify(aBimap.GetValue("Hello")))
	fmt.Println(stringify(aBimap.GetKey(1)))
	fmt.Println(stringify(aBimap.GetValue("guy")))
	fmt.Println(stringify(aBimap.GetKey(-1)))

	aBimap.Put("guy", 2)
	assert.Equal(2, aBimap.Size())

	aBimap.RemoveKey("Hello")
	assert.Equal(1, aBimap.Size())

	aBimap.RemoveValue(2)
	assert.Equal(0, aBimap.Size())
}

func TestDuplicatedEntries(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()

	aBimap.Put("Hello", 1)
	assert.Equal(1, aBimap.Size())
	v, _ := aBimap.GetValue("Hello")
	assert.Equal(1, v)

	aBimap.Put("Hello", 2)
	assert.Equal(1, aBimap.Size())
	v, _ = aBimap.GetValue("Hello")
	assert.Equal(2, v)
}

func TestInverse(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()
	aBimap.Put("Hello", 1)
	aBimap.Put("guy", 2)
	assert.Equal(2, aBimap.Size())

	iBimap := aBimap.Inverse()
	v, _ := iBimap.GetValue(1)
	assert.Equal("Hello", v)
	v, _ = iBimap.GetValue(2)
	assert.Equal("guy", v)
	assert.Equal(2, iBimap.Size())

	iBimap.RemoveKey(1)
	assert.Equal(1, iBimap.Size())
	assert.Equal(2, aBimap.Size())
	v, _ = iBimap.GetValue(2)
	assert.Equal("guy", v)
}
