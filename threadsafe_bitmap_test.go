package bitmap

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBitConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var size uint64 = 50

	bm := newThreadSafeBitmap(size)

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
