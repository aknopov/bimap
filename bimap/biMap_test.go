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
	assert.False(aBimap.ContainsKey("there!"))

	fmt.Println(stringify(aBimap.GetValue("Hello")))
	fmt.Println(stringify(aBimap.GetKey(1)))
	fmt.Println(stringify(aBimap.GetValue("there!")))
	fmt.Println(stringify(aBimap.GetKey(-1)))

	aBimap.Put("there!", 2)
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
}

func TestKeysValues(t *testing.T) {
	assert := assert.New(t)

	aBimap := NewBiMap[string, int]()
	aBimap.Put("Hello", 1)
	aBimap.Put("there!", 2)

	assert.Equal(map[string]Void{"Hello": Null, "there!": Null}, aBimap.Keys())
	assert.Equal(map[int]Void{1: Null, 2: Null}, aBimap.Values())
}
