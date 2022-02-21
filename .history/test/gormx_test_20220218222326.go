package test

import (
	"gdemo1/gormx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGormx(t *testing.T) {

	cfg := gormx.Config{
		DBType: "mysql",
		DSN:    "fage:fage501526~@tcp(127.0.0.1:3306)",
	}
	db, err := cfg.New()
	assert.Nil(err)
}
