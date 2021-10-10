package bitmap

type Bitmap interface {
	SetBIT(offset uint64, v bool) bool
	GetBit(offset uint64) bool
	Count() uint64
}
