package repository

import (
	"os"
	"testing"

	"github.com/carlosvinicius3258/go-transaction/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDb_Insert(t * testing.T){
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2, "approved", "")
	assert.Nil(t, err)

}