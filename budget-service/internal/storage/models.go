// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

func (e *TransactionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionType(s)
	case string:
		*e = TransactionType(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionType: %T", src)
	}
	return nil
}

type NullTransactionType struct {
	TransactionType TransactionType
	Valid           bool // Valid is true if TransactionType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionType) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionType), nil
}

type Budget struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CategoryID uuid.NullUUID
	Amount     int64
	Spent      sql.NullInt64
	Currency   string
}

type Category struct {
	ID   uuid.UUID
	Name string
}

type Transaction struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Type       TransactionType
	Amount     int64
	Currency   string
	CategoryID uuid.NullUUID
	Date       sql.NullTime
}
