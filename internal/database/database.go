package database

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db  *sql.DB
	mtx sync.Mutex
}

func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// func (m *Database) CreateTables() error {
// 	m.mtx.Lock()
// 	defer m.mtx.Unlock()

// 	_, err := m.db.Exec("CREATE TABLE IF NOT EXISTS addresses(address PRIMARY KEY)")

// 	return err
// }

func (m *Database) Close() error {
	return m.db.Close()
}

func (m *Database) SelectAddress(address string) (bool, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	rows, err := m.db.Query("SELECT address FROM addresses WHERE address = ?", address)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	result := false
	error_ := new(error)

	for rows.Next() {
		var address string

		err = rows.Scan(&address)

		if err != nil {
			result = false
			error_ = &err
		}

		result = true

	}

	return result, *error_

}
