//go:build unit

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	m, err := NewBook(uint64(1), "Name 1")
	assert.Nil(t, err)
	assert.Equal(t, m.ID, uint64(1))
}

func TestBookValidate(t *testing.T) {
	testCases := []struct {
		id   uint64
		name string
		err  error
	}{
		{
			id:   1,
			name: "Name 1",
			err:  nil,
		},
		{
			id:   2,
			name: "",
			err:  ErrInvalidModel,
		},
	}
	for _, tc := range testCases {
		_, err := NewBook(tc.id, tc.name)
		assert.Equal(t, err, tc.err)
	}
}
