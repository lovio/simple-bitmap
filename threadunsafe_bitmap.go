package bitmap

import "math"

// restrict bitmap size to 512MB
const MaxBitmapSize uint64 = math.MaxUint32

type ThreadUnsafeBitmap struct {
	data []byte
	size uint64 // support larger MaxBitmapSize for future
}

func newThreadUnsafeBitmap(size uint64) *ThreadUnsafeBitmap {
	if size < 1 || size > MaxBitmapSize {
		size = MaxBitmapSize
	}
	// set data length to 1 if size <= 7
	length := size & 7 // size % 8
	if length != 0 {
		length = 1
	}
	return &ThreadUnsafeBitmap{
		data: make([]byte, (size>>3)+length),
		size: size,
	}
}

func (b *ThreadUnsafeBitmap) SetBit(offset uint64, v bool) bool {
	if offset > b.size {
		return false
	}
	index, bitPosition := getBitmapPosition(offset)
	if v {
		b.data[index] |= 1 << bitPosition // set bitPostion to true
	} else {
		b.data[index] &^= 1 << bitPosition // bit clear by bitPosition
	}
	return true
}

func (b *ThreadUnsafeBitmap) GetBit(offset uint64) bool {
	if offset > b.size {
		return false
	}
	index, bitPosition := getBitmapPosition(offset)
	return (b.data[index]>>bitPosition)&1 != 0
}

func (b *ThreadUnsafeBitmap) Count() uint64 {
	return b.size
}
