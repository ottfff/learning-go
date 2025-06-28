package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceMake(t *testing.T) {
	var s1 []int
	s2 := make([]int, 4, 7)
	s3 := []int{1, 2, 3, 4, 5}

	assert.Nil(t, s1)
	assert.Equal(t, 0, len(s1))
	assert.Equal(t, 0, cap(s1))

	assert.Equal(t, 4, len(s2))
	assert.Equal(t, 7, cap(s2))
	assert.Equal(t, []int{0, 0, 0, 0}, s2)

	assert.Equal(t, 5, len(s3))
	assert.Equal(t, 5, cap(s3))
}

func TestSliceAppend(t *testing.T) {
	var s []int
	s2 := append(s, 5)

	assert.Equal(t, []int{5}, s2)
	assert.Equal(t, 1, len(s2))
	assert.Equal(t, 1, cap(s2))

	assert.Nil(t, s)
	assert.Equal(t, 0, len(s))
	assert.Equal(t, 0, cap(s))

	s2 = append(s2, 10)
	assert.Equal(t, []int{5, 10}, s2)
	assert.Equal(t, 2, len(s2))
	assert.Equal(t, 2, cap(s2))

	s2 = append(s2, 10)
	assert.Equal(t, []int{5, 10, 10}, s2)
	assert.Equal(t, 3, len(s2))
	assert.Equal(t, 4, cap(s2))

	s2 = append(s2, 10)
	assert.Equal(t, []int{5, 10, 10, 10}, s2)
	assert.Equal(t, 4, len(s2))
	assert.Equal(t, 4, cap(s2))

	s2 = append(s2, 10)
	assert.Equal(t, []int{5, 10, 10, 10, 10}, s2)
	assert.Equal(t, 5, len(s2))
	assert.Equal(t, 8, cap(s2))

	s3 := make([]int, 511)
	assert.Equal(t, 511, len(s3))
	assert.Equal(t, 511, cap(s3))

	s3 = append(s3, 0)
	assert.Equal(t, 512, len(s3))
	assert.Equal(t, 848, cap(s3))

}
