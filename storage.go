package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAllAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres password=postgres dbname=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	err := s.createAccountTable()
	return err
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS ACCOUNT(
		ID SERIAL PRIMARY KEY,
		FIRSTNAME VARCHAR(50),
		LASTNAME VARCHAR(50),
		NUMBER SERIAL,
		BALANCE SERIAL,
		CREATED_AT TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(req *Account) error {
	query := `insert into account 
	(firstname, lastname, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(
		query,
		req.FirstName,
		req.LastName,
		req.Number,
		req.Balance,
		req.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {
	query := `select * from account`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	response := []*Account{}
	for rows.Next() {
		t := new(Account)
		err := rows.Scan(
			&t.ID,
			&t.FirstName,
			&t.LastName,
			&t.Number,
			&t.Balance,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		response = append(response, t)
	}

	return response, nil
}
