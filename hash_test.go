package safemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genHasher(t *testing.T) {
	hasher := genHasher[string]()
	assert.Equal(t, hasher("499"), hasher("499"))
	assert.NotEqual(t, hasher("499"), hasher("500"))

	hasher1 := genHasher[int]()
	assert.Equal(t, hasher1(499), hasher1(499))
	assert.NotEqual(t, hasher1(499), hasher1(500))

	hasher2 := genHasher[uint]()
	assert.Equal(t, hasher2(499), hasher2(499))
	assert.NotEqual(t, hasher2(499), hasher2(500))

	hasher3 := genHasher[uintptr]()
	assert.Equal(t, hasher3(128), hasher3(128))
	assert.NotEqual(t, hasher3(128), hasher3(129))

	hasher4 := genHasher[byte]()
	assert.Equal(t, hasher4(128), hasher4(128))
	assert.NotEqual(t, hasher4(128), hasher4(129))

	hasher5 := genHasher[int8]()
	assert.Equal(t, hasher5(100), hasher5(100))
	assert.NotEqual(t, hasher5(100), hasher5(101))

	hasher6 := genHasher[uint8]()
	assert.Equal(t, hasher6(128), hasher6(128))
	assert.NotEqual(t, hasher6(128), hasher6(129))

	hasher7 := genHasher[int16]()
	assert.Equal(t, hasher7(100), hasher7(100))
	assert.NotEqual(t, hasher7(100), hasher7(101))

	hasher8 := genHasher[uint16]()
	assert.Equal(t, hasher8(128), hasher8(128))
	assert.NotEqual(t, hasher8(128), hasher8(129))

	hasher9 := genHasher[int32]()
	assert.Equal(t, hasher9(100), hasher9(100))
	assert.NotEqual(t, hasher9(100), hasher9(101))

	hasher10 := genHasher[uint32]()
	assert.Equal(t, hasher10(128), hasher10(128))
	assert.NotEqual(t, hasher10(128), hasher10(129))

	hasher11 := genHasher[float32]()
	assert.Equal(t, hasher11(128), hasher11(128))
	assert.NotEqual(t, hasher11(128), hasher11(129))

	hasher12 := genHasher[int64]()
	assert.Equal(t, hasher12(100), hasher12(100))
	assert.NotEqual(t, hasher12(100), hasher12(101))

	hasher13 := genHasher[uint64]()
	assert.Equal(t, hasher13(128), hasher13(128))
	assert.NotEqual(t, hasher13(128), hasher13(129))

	hasher14 := genHasher[float64]()
	assert.Equal(t, hasher14(128), hasher14(128))
	assert.NotEqual(t, hasher14(128), hasher14(129))

	hasher15 := genHasher[complex64]()
	assert.Equal(t, hasher15(complex(1, 2)), hasher15(complex(1, 2)))
	assert.NotEqual(t, hasher15(complex(1, 2)), hasher15(complex(1, 3)))

	hasher16 := genHasher[complex128]()
	assert.Equal(t, hasher16(complex(1, 2)), hasher16(complex(1, 2)))
	assert.NotEqual(t, hasher16(complex(1, 2)), hasher16(complex(1, 3)))

	// pointer to structs
	type S struct {
		a int
	}

	hasher17 := genHasher[*S]()
	s1 := &S{
		a: 1,
	}
	s2 := &S{
		a: 2,
	}
	assert.Equal(t, hasher17(s1), hasher17(s1))
	assert.NotEqual(t, hasher17(s1), hasher17(s2))

	// channel
	hasher18 := genHasher[chan int]()
	ch1 := make(chan int)
	ch2 := make(chan int)
	assert.Equal(t, hasher18(ch1), hasher18(ch1))
	assert.NotEqual(t, hasher18(ch1), hasher18(ch2))

	// https://github.com/golang/go/issues/51338
	//
	// // interface is not supported
	// hasher19 := genHasher[any]()
	// var i1 any = 1
	// var i2 any = []int{1}
	// assert.Equal(t, hasher19(i1), hasher19(i1))
	// assert.NotEqual(t, hasher19(i1), hasher19(i2))
	// // but builtin map supports it.
	// m := make(map[any]int)
	// m[i1] = 1
	// m[i2] = 2
	// assert.Equal(t, m[i1], m[i1])
	// assert.NotEqual(t, m[i1], m[i2])

	// map, slice, fubnction are not comparable
}
