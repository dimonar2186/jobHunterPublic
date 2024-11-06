package hrcontact

//Basic structure description of hrcontact and basic methods are provided here.
//Important: Save, Update, Delete, StatusUpdate, GetAllHrcontactsOfHrmanager require an opened transaction to be executed

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	hRContactTable string = "public.hrcontacts"
)

type HRContact struct {
	Id           string
	Value        string
	Messenger    string
	Preffered    bool
	IsDeleted    bool
	HRManagerId  string
	CreationDate time.Time
	UpdateDate   time.Time
}

// applicationStage constructor
func NewHRContact(
	id, value, messenger, hRManagerId string,
	preffered, isDeleted bool,
	creationDate, updateDate time.Time) *HRContact {
	return &HRContact{
		Id:           id,
		Value:        value,
		Messenger:    messenger,
		Preffered:    preffered,
		IsDeleted:    isDeleted,
		HRManagerId:  hRManagerId,
		CreationDate: creationDate,
		UpdateDate:   updateDate}
}

// HRContact select from DB record
func (hRContact *HRContact) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, value, messenger, preferred, is_deleted, hr_manager, created_at, updated_at FROM $1 where id = $2", hRContactTable, id) //read the HRContact from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(hRContact.Id, hRContact.Value, hRContact.Messenger, hRContact.Preffered, hRContact.IsDeleted, hRContact.HRManagerId,
			hRContact.CreationDate, hRContact.UpdateDate); err != nil {
			return err
		} else {
			log.Println("Failed to extract HRContact with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("HRContact with id=%q not found", id)
}

// Saving HRContact to DB
func (hRContact *HRContact) Save(tx *sql.Tx) error {
	//form the query
	hRContactDBCols := "(value, messenger, preferred, is_deleted, hr_manager, created_at, updated_at)"
	hRContactInsertParameters := "($1, $2, $3, $4, $5, $6, $7)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		hRContactTable, hRContactDBCols, hRContactInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		hRContact.Value, hRContact.Messenger, hRContact.Preffered, hRContact.IsDeleted, hRContact.HRManagerId,
		time.Now(), time.Now()).Scan(&id); err != nil {
		return fmt.Errorf("failed to save HRContact %w", err)
	}
	hRContact.Id = id
	return nil
}

// Deleting the HRContact from DB
func (hRContact *HRContact) Delete(tx *sql.Tx) error {
	//form deleting query and execute it
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = %s", hRContactTable, hRContact.Id)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("failed to delete the HRContact %w", err)
	}
	return nil
}

func (hRContact *HRContact) StatusUpdate(status string, tx *sql.Tx) error {
	//There is no field called status.
	return nil
}

func (hRContact *HRContact) Update(tx *sql.Tx) error {
	//Form and execute update query
	hRContactDBCols := "(value, messenger, preferred, is_deleted, hr_manager, created_at, updated_at)"
	hRContactInsertParameters := "($1, $2, $3, $4, $5, $6, $7)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", hRContactTable, hRContactDBCols, hRContactInsertParameters, hRContact.Id)
	if _, err := tx.Exec(query,
		hRContact.Value, hRContact.Messenger, hRContact.Preffered, hRContact.IsDeleted, hRContact.HRManagerId,
		hRContact.CreationDate, time.Now()); err != nil {
		errorMessage := "failed to update the HRContact with id=" + hRContact.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Getter for Id
func (hRContact *HRContact) GetId() string {
	return hRContact.Id
}
