package store

import (
	"database/sql"
)

func Open() (*sql.DB, error) {

	db, err := sql.Open("pgx")

}
