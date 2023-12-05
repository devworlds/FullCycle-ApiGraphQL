package database

import (
	"database/sql"
)

type Token struct {
	db          *sql.DB
	Address     string
	Name        string
	Description string
}

func NewToken(db *sql.DB) *Token {
	return &Token{db: db}
}

func (t *Token) Create(address, name, description string) (Token, error) {
	_, err := t.db.Exec("INSERT INTO tokens (address, name, description) VALUES ($1, $2, $3)", address, name, description)
	if err != nil {
		return Token{}, err
	}
	return Token{Address: address, Name: name, Description: description}, nil
}

func (t *Token) FindAll() ([]Token, error) {
	rows, err := t.db.Query("SELECT address, name, description FROM tokens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []Token{}
	for rows.Next() {
		var address, name, description string
		if err := rows.Scan(&address, &name, &description); err != nil {
			return nil, err
		}
		tokens = append(tokens, Token{Address: address, Name: name, Description: description})
	}
	return tokens, nil
}
