package bitmap

import "sync"

type ThreadSafeBitmap struct {
	unsafeBitmap Bitmap
	sync.RWMutex
}

func newThreadSafeBitmap(size uint64) *ThreadSafeBitmap {
	return &ThreadSafeBitmap{
		unsafeBitmap: newThreadUnsafeBitmap(size),
	}
}

func (b *ThreadSafeBitmap) SetBit(offset uint64, v bool) bool {
	b.Lock()
	defer b.Unlock()
	return b.unsafeBitmap.SetBit(offset, v)
}

func (b *ThreadSafeBitmap) GetBit(offset uint64) bool {
	b.RLock()
	defer b.RUnlock()
	return b.unsafeBitmap.GetBit(offset)
}

func (b *ThreadSafeBitmap) Count() uint64 {
	return b.unsafeBitmap.Count()
}
