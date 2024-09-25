package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil)

	// assert for not nil (good when you expect something)
	i := new(int)
	if assert.NotNil(t, i) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", i)
	}
}
