package test

import (
	"gdemo1/casbin"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasbin(t *testing.T) {

	bl := casbin.New()
	assert.True(t, bl)
}
