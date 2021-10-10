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

func TestUnion(t *testing.T) {
	b := NewBitmap(100)
	other := NewBitmap(200)

	b.SetBit(1, true)
	other.SetBit(100, true)

	bm := b.Union(other)
	assert.Equal(t, uint64(200), bm.Count())

	assert.True(t, bm.GetBit(1))
	assert.True(t, bm.GetBit(100))
	assert.False(t, bm.GetBit(2))
	assert.False(t, bm.GetBit(200))
}

func TestIntersect(t *testing.T) {
	{
		t.Log("b is smaller than other")
		b := NewBitmap(100)
		other := NewBitmap(200)
		b.SetBit(1, true)
		b.SetBit(100, true)
		other.SetBit(1, true)
		other.SetBit(200, true)

		bm := b.Intersect(other)
		assert.Equal(t, b.Count(), bm.Count())

		assert.True(t, bm.GetBit(1))
		assert.False(t, bm.GetBit(100))
		assert.False(t, bm.GetBit(200))
	}

	{
		t.Log("b is larger than other")
		b := NewBitmap(200)
		other := NewBitmap(100)
		b.SetBit(1, true)
		b.SetBit(200, true)
		other.SetBit(1, true)
		other.SetBit(100, true)

		bm := b.Intersect(other)
		assert.Equal(t, b.Count(), bm.Count())

		assert.True(t, bm.GetBit(1))
		assert.False(t, bm.GetBit(100))
		assert.False(t, bm.GetBit(200))
	}
}

func TestDifference(t *testing.T) {
	b := NewBitmap(100)
	other := NewBitmap(100)
	b.SetBit(1, true)
	b.SetBit(100, true)
	other.SetBit(1, true)
	other.SetBit(2, true)

	bm := b.Difference(other)
	assert.Equal(t, b.Count(), bm.Count())

	assert.False(t, bm.GetBit(1))
	assert.False(t, bm.GetBit(2))
	assert.True(t, bm.GetBit(100))

}
