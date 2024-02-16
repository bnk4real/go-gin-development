package database

import (
	"fmt"
	"testing"
)

func Test_Connection(t *testing.T) {

	db, err := NewPostgresqlDatabase()
	if err != nil {
		t.Error(db)
	}

	fmt.Println("Connected Successfully")
}
