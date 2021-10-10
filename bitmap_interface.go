package bitmap

type Bitmap interface {
	SetBit(offset uint64, v bool) bool
	GetBit(offset uint64) bool
	Count() uint64
}

func NewThreadUnsafeBitmap(size uint64) Bitmap {
	return newThreadUnsafeBitmap(size)
}
