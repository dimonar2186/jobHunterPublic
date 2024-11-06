package entity

import (
	"database/sql"
)

// Entity — interface of basic DB operatins
type Entity interface {
	// Saving new entity in DB
	Save(tx *sql.Tx) error

	// FindById searches the existing entyti by id.
	FindById(id string, tx *sql.Tx) error

	// Delete the existing entity chosen by id.
	Delete(tx *sql.Tx) error

	//Update status of the existing entity
	StatusUpdate(status string, tx *sql.Tx) error

	//Update the existing entity
	Update(tx *sql.Tx) error

	//Getter for Id
	GetId() string
}
