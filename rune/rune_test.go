package rune

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	s := "Hello World!"
	bytes := []byte(s)
	assert.Equal(t, 12, len(bytes))
	assert.Equal(t, s, string(bytes))

	rus_s := "Привет!"
	bytes = []byte(rus_s)
	assert.Equal(t, 13, len(bytes))
	assert.Equal(t, rus_s, string(bytes))

	emoji := "😀"
	assert.Equal(t, 4, len(emoji))
	assert.Equal(t, 1, len([]rune(emoji)))
	assert.Equal(t, 1, utf8.RuneCountInString(emoji))

}
