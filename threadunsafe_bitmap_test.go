package bitmap

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadUnsafeBitmap(t *testing.T) {
	{
		bm := NewThreadUnsafeBitmap(math.MaxUint64)
		assert.Equal(t, MaxBitmapSize, bm.Count(), "restrict max bitmap size to 512MB")
	}

	{
		bm := NewThreadUnsafeBitmap(50)
		assert.Equal(t, uint64(50), bm.Count())
		assert.False(t, bm.SetBit(60, true), "offset is larger than bitmap size")

		assert.True(t, bm.SetBit(30, true))
		assert.True(t, bm.GetBit(30))
		assert.False(t, bm.GetBit(29))

		assert.True(t, bm.SetBit(30, false))
		assert.False(t, bm.GetBit(30))
	}
}
