package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB(source string) (*DB, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, fmt.Errorf("cannot open connect to db %v: %w", source, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db %v: %w", source, err)
	}

	return &DB{db}, nil
}

func (db *DB) IsNoRowsErr(err error) bool {
	return err == sql.ErrNoRows
}
