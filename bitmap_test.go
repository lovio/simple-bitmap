package bitmap

import (
	"math"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBit(t *testing.T) {
	{
		bm := NewBitmap(math.MaxUint64)
		assert.Equal(t, MaxBitmapSize, bm.Count(), "restrict max bitmap size to 512MB")
	}

	{
		bm := NewBitmap(50)
		assert.Equal(t, uint64(50), bm.Count())
		assert.False(t, bm.SetBit(60, true), "offset is larger than bitmap size")

		assert.True(t, bm.SetBit(30, true))
		assert.True(t, bm.GetBit(30))
		assert.False(t, bm.GetBit(29))

		assert.True(t, bm.SetBit(30, false))
		assert.False(t, bm.GetBit(30))

		assert.False(t, bm.GetBit(51), "offset larger than size should be false")
	}
}

func TestSetBitConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var size uint64 = 100000000

	bm := NewBitmap(size)

	ints := 50

	var wg sync.WaitGroup
	wg.Add(ints)
	for i := 0; i < ints; i++ {
		go func(i int) {
			assert.True(t, bm.SetBit(uint64(i), true))
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < ints; i++ {
		assert.True(t, bm.GetBit((uint64(i))))
	}

	assert.Equal(t, size, bm.Count())

}
