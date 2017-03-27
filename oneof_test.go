package flagvals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOfStringValid(t *testing.T) {
	assert := assert.New(t)
	f := NewOneOfString("one", "two")
	err := f.Set("two")
	assert.Nil(err)
	assert.Equal(f.String(), "two")
}

func TestOneOfStringInvalid(t *testing.T) {
	assert := assert.New(t)
	f := NewOneOfString("one", "two")
	err := f.Set("three")
	assert.NotNil(err)
	assert.Equal(f.String(), "")
}
