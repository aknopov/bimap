package bimap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiMapBasics(t *testing.T) {
	aBimap := NewBiMap[string, int]()
	assert.Equal(t, 0, aBimap.Size())

	aBimap.Put("Hello", 1)
	assert.Equal(t, 1, aBimap.Size())

	//TODO How to write "stringify()"?
	if val, ok := aBimap.GetValue("Hello"); ok {
		fmt.Printf("%+v\n", val)
	} else {
		fmt.Println("<nil>")
	}
	if key, ok := aBimap.GetKey(1); ok {
		fmt.Printf("%+v\n", key)
	} else {
		fmt.Println("nothing!")
	}
	if val, ok := aBimap.GetValue("guy"); ok {
		fmt.Printf("%+v\n", val)
	} else {
		fmt.Println("<nil>")
	}
	if key, ok := aBimap.GetKey(-1); ok {
		fmt.Printf("%+v\n", key)
	} else {
		fmt.Println("nothing!")
	}

	aBimap.Put("guy", 2)
	assert.Equal(t, 2, aBimap.Size())

	aBimap.RemoveKey("Hello")
	assert.Equal(t, 1, aBimap.Size())

	aBimap.RemoveValue(2)
	assert.Equal(t, 0, aBimap.Size())
}
