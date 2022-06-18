package main_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	_db "github.com/questizen/core-system/app"
	"github.com/stretchr/testify/assert"
)

func TestDbConnection(t *testing.T) {
	db, err := _db.SetupDb()

	assert.NoError(t, err)
	defer db.Close()
}
