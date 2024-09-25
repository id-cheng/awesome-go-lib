package testify

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequire(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello World"

	require.Equal(t, a, b, "The two words should be the same.")

}
