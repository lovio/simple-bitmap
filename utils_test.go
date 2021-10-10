package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBitmapPosition(t *testing.T) {
	testCases := []struct {
		desc        string
		offset      uint64
		index       uint64
		bitPosition uint64
	}{
		{
			desc:        "",
			offset:      0,
			index:       0,
			bitPosition: 0,
		},
		{
			desc:        "",
			offset:      31,
			index:       0,
			bitPosition: 31,
		},
		{
			desc:        "",
			offset:      32,
			index:       1,
			bitPosition: 0,
		},
		{
			desc:        "",
			offset:      64,
			index:       2,
			bitPosition: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			index, bitPosition := getBitmapPosition(tC.offset)
			assert.Equal(t, tC.index, index)
			assert.Equal(t, tC.bitPosition, bitPosition)
		})
	}
}
