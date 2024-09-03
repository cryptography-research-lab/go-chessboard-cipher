package chessboard_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	assert.Equal(t, 'c', ToLowerCase('C'))
}

func TestToUpperCase(t *testing.T) {
	assert.Equal(t, 'C', ToUpperCase('c'))
}
