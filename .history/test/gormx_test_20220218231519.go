package test

import (
	"fmt"
	"gdemo1/gormx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGormx(t *testing.T) {

	cfg := gormx.Config{
		DBType: "mysql",
		DSN:    "fage:Fage501526~@/mytest?charset=utf8&parseTime=True&loc=Local",
	}
	db, err := cfg.New()
	assert.Nil(t, err)
	query := fmt.Sprintf("create database if not exists 'mytest' default character set ='utf8mb4';")
	_, err = db.Exec(query).DB()
	assert.Nil(err)
}
