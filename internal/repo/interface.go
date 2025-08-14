package repo

import (
	"database/sql"
)

type Dao interface {
	Reset() error
	AddPhoneNumber(phone string) error
	GetAllPhoneNumbers() ([]string, error)
}

func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}

type dao struct {
	db *sql.DB
}

func (d *dao) AddPhoneNumber(s string) error {
	query := "INSERT INTO numbers (num) VALUES (?)"
	_, err := d.db.Exec(query, s)

	return err

}

func (d *dao) GetAllPhoneNumbers() ([]string, error) {
	result := []string{}
	query := "SELECT  num FROM numbers"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pn string
		if err := rows.Scan(&pn); err != nil {
			return nil, err
		}
		result = append(result, pn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *dao) Reset() error {
	query := "TRUNCATE numbers"
	_,  err := d.db.Exec(query)

	return err
}
