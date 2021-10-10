package bitmap

import (
	"math"
	"sync/atomic"
)

// restrict bitmap size to 512MB
const MaxBitmapSize uint64 = math.MaxUint32

type Bitmap interface {
	SetBit(offset uint64, v bool) bool
	GetBit(offset uint64) bool
	Count() uint64
}

func NewBitmap(size uint64) Bitmap {
	return newThreadsafeBitmap(size)
}

type threadsafeBitmap struct {
	data []uint32
	size uint64 // support larger MaxBitmapSize for future
}

func newThreadsafeBitmap(size uint64) *threadsafeBitmap {
	if size < 1 || size > MaxBitmapSize {
		size = MaxBitmapSize
	}
	// set data length to 1 if size <= 31
	length := size & 31 // size % 31
	if length != 0 {
		length = 1
	}
	return &threadsafeBitmap{
		data: make([]uint32, (size>>5)+length),
		size: size,
	}
}

func (b *threadsafeBitmap) SetBit(offset uint64, v bool) bool {
	if offset > b.size {
		return false
	}
	index, bitPosition := getBitmapPosition(offset)
	p := &(b.data[index])
	var oldValue, newValue uint32
	for {
		oldValue = atomic.LoadUint32(p)
		if v {
			newValue = oldValue | 1<<bitPosition
		} else {
			newValue = oldValue &^ 1 << uint32(bitPosition)
		}
		if atomic.CompareAndSwapUint32(p, oldValue, newValue) {
			return true
		}
	}
}

func (b *threadsafeBitmap) GetBit(offset uint64) bool {
	if offset > b.size {
		return false
	}
	index, bitPosition := getBitmapPosition(offset)
	return (b.data[index]>>bitPosition)&1 != 0
}

func (b *threadsafeBitmap) Count() uint64 {
	return b.size
}
