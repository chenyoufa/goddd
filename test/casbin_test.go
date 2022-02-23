package test

import (
	"gdemo1/casbin"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasbin(t *testing.T) {

	// bl := casbin.ACL()
	// assert.True(t, bl)
	bl2 := casbin.RBAC()
	assert.True(t, bl2)
}
