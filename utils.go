package bitmap

// bit operation
// GET slices index: offset >> 3 equals to offset/8
// GET bit position index: offset & 7  equals to offset%8
func getBitmapPosition(offset uint64) (index, bitPosition uint64) {
	return offset >> 3, offset & 7
}
