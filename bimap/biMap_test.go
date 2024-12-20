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

func assertPanic(t *testing.T, msg string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(msg)
		}
	}()
	f()
}

func TestNewBiMap(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMapEx[string, int](5)

	assert.Equal(0, aBimap.Size())
}

func TestBiMapBasics(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()

	aBimap.Put("Hello", 1)
	assert.Equal(1, aBimap.Size())

	fmt.Println(stringify(aBimap.GetValue("Hello")))
	fmt.Println(stringify(aBimap.GetKey(1)))
	fmt.Println(stringify(aBimap.GetValue("there!")))
	fmt.Println(stringify(aBimap.GetKey(-1)))

	assert.True(aBimap.ContainsKey("Hello"))
	assert.True(aBimap.ContainsValue(1))
	assert.False(aBimap.ContainsKey("there!"))

	aBimap.Put("there!", 2)
	assert.Equal(2, aBimap.Size())

	aBimap.RemoveKey("Hello")
	assert.Equal(1, aBimap.Size())
	idx, ok := aBimap.GetValue("there!")
	assert.True(ok)
	assert.Equal(2, idx)
	_, ok = aBimap.GetValue("Hello")
	assert.False(ok)

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
	aBimap.Put("there!", 2)
	assert.Equal(2, aBimap.Size())

	iBimap := aBimap.Inverse()
	v, _ := iBimap.GetValue(1)
	assert.Equal("Hello", v)
	v, _ = iBimap.GetValue(2)
	assert.Equal("there!", v)
	assert.Equal(2, iBimap.Size())

	iBimap.RemoveKey(1)
	assert.Equal(1, iBimap.Size())
	assert.Equal(2, aBimap.Size())
	v, _ = iBimap.GetValue(2)
	assert.Equal("there!", v)
}

func TestEquals(t *testing.T) {
	assert := assert.New(t)

	bimap1 := NewBiMap[string, int]()
	bimap2 := NewBiMap[string, int]()
	assert.True(bimap1.Equals(bimap2))
	assert.True(bimap2.Equals(bimap1))

	bimap1.Put("Hello", 1)
	bimap1.Put("there!", 2)
	assert.False(bimap1.Equals(bimap2))
	assert.False(bimap2.Equals(bimap1))

	bimap2.Put("Hello", 1)
	assert.False(bimap1.Equals(bimap2))
	assert.False(bimap2.Equals(bimap1))
	bimap2.Put("there!", 2)
	assert.True(bimap1.Equals(bimap2))
	assert.True(bimap2.Equals(bimap1))
}

func TestPutAll(t *testing.T) {
	assert := assert.New(t)

	bimap1 := NewBiMap[string, int]()
	bimap2 := NewBiMap[string, int]()
	bimap1.Put("Hello", 1)
	bimap1.Put("there!", 2)

	assert.Equal(0, bimap2.Size())
	bimap2.PutAll(bimap1)
	assert.Equal(2, bimap2.Size())

	// UC duplicated keys/values
}

func TestKeysValues(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()
	aBimap.Put("Hello", 1)
	aBimap.Put("there!", 2)

	assert.Equal([]string{"Hello", "there!"}, aBimap.Keys())
	assert.Equal([]int{1, 2}, aBimap.Values())
}

func TestIterator(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()
	message := []string{"Hi ", "there!", " Here ", " goes ", " the ", "test."}
	for i, v := range message {
		aBimap.Put(v, i+1)
	}

	it := aBimap.Iterator()
	i := 0
	for it.HasNext() {
		k, v := it.Next()
		assert.Equal(message[i], k)
		assert.Equal(i+1, v)
		i++
	}

	assertPanic(t, "No panic?", func() { it.Next() })
}
