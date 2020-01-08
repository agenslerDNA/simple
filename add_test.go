package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/agenslerDNA/simple"
)

func TestAdd(t *testing.T) {
	ans := Add(1, 1)
	assert.Equal(t, 1, ans)
}
