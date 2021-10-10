package bitmap

// bit operation
// GET slices index: offset >> 5 equals to offset/32
// GET bit position index: offset & 31  equals to offset%32
func getBitmapPosition(offset uint64) (index, bitPosition uint64) {
	return offset >> 5, offset & 31
}
