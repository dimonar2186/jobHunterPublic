package hrmanager

//Basic structure description of HRManager and basic methods are provided here.
//Important: Save, Update, Delete, GetAllContactsOfHRManager require an opened transaction to be executed

import (
	"JobHunterBackService/dbentity/hrcontact"
	"JobHunterBackService/dbentity/vacancyhrmanagerlink"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	hRManagerTable string = "public.hrmanagers"
	hRContactTable string = "public.hrcontacts"
)

type HRManager struct {
	Id           string
	FirstName    string
	LastName     string
	SecondName   string
	EmployerId   string
	CreationDate time.Time
}

// HRManager constructor
func NewHRManager(
	id, firstName, lastName, secondName, employerId string,
	creationDate time.Time) *HRManager {
	return &HRManager{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		SecondName:   secondName,
		EmployerId:   employerId,
		CreationDate: creationDate}
}

// HRManager select from DB record
func (hRManager *HRManager) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, first_name, last_name, second_name, employer, created_at FROM $1 where id = $2", hRManagerTable, id) //read the HRManager from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(hRManager.Id, hRManager.FirstName, hRManager.LastName, hRManager.SecondName,
			hRManager.EmployerId, hRManager.CreationDate); err != nil {
			return err
		} else {
			log.Println("Failed to extract HRManager with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("HRManager with id=%q not found", id)
}

// Saving HRManager to DB
func (hRManager *HRManager) Save(tx *sql.Tx) error {
	//form the query
	hRManagerDBCols := "(first_name, last_name, second_name, employer, created_at)"
	hRManagerInsertParameters := "($1, $2, $3, $4, $5)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		hRManagerTable, hRManagerDBCols, hRManagerInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		hRManager.FirstName, hRManager.LastName, hRManager.SecondName,
		hRManager.EmployerId, time.Now()).Scan(&id); err != nil {
		return fmt.Errorf("failed to save HRManager %w", err)
	}
	hRManager.Id = id
	return nil
}

// Deleting the HRManager from DB
func (hRManager *HRManager) Delete(tx *sql.Tx) error {
	//find all contacs of HRManager and then delete all of them
	var allHRContacts []*hrcontact.HRContact
	var err error
	if allHRContacts, err = hRManager.GetAllContactsOfHRManager(tx); err != nil {
		return fmt.Errorf("during HRManager deletion %w", err)
	}
	if len(allHRContacts) > 0 {
		for contactNumber := 0; contactNumber < len(allHRContacts); contactNumber++ {
			allHRContacts[contactNumber].Delete(tx)
		}
		log.Println("HRContact.id: ", &hRManager.Id, " contacts deleted")
	}

	//delete all links between hrmanager and vacancies
	if err = vacancyhrmanagerlink.Delete(hRManager, tx); err != nil {
		return err
	}

	//form deleting query and execute it
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = %s", hRManagerTable, hRManager.Id)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("failed to delete the HRManger %w", err)
	}
	return nil
}

func (hRManager *HRManager) StatusUpdate(status string, tx *sql.Tx) error {
	//It is unneccessary function, because there is no such field
	return nil
}

func (hRManager *HRManager) Update(tx *sql.Tx) error {
	//Form and execute update query
	hRManagerDBCols := "(first_name, last_name, second_name, employer)"
	hRManagerInsertParameters := "($1, $2, $3, $4)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", hRManagerTable, hRManagerDBCols, hRManagerInsertParameters, hRManager.Id)
	if _, err := tx.Exec(query,
		hRManager.FirstName, hRManager.LastName, hRManager.SecondName, hRManager.EmployerId); err != nil {
		errorMessage := "failed to update the HRManager with id=" + hRManager.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Getter for Id
func (hRManager *HRManager) GetId() string {
	return hRManager.Id
}

// form full list of HRContacts for the HRManager
func (hRManager *HRManager) GetAllContactsOfHRManager(tx *sql.Tx) (allHRContacts []*hrcontact.HRContact, err error) {

	rows, err := tx.Query("SELECT id, value, messenger, preferred, is_deleted, hr_manager, created_at, updated_at FROM $1 where hr_manager = $2", hRContactTable, hRManager.Id) //read the HRContacts list from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty temporary HRContact
		tempHRContact := hrcontact.NewHRContact("", "", "", "", false, false, time.Time{}, time.Time{})
		//fulfill this temp HRContact with row data and then append it into the slice of HRContacts
		if err := rows.Scan(tempHRContact.Id, tempHRContact.Value, tempHRContact.Messenger, tempHRContact.Preffered, tempHRContact.IsDeleted, tempHRContact.HRManagerId,
			tempHRContact.CreationDate, tempHRContact.UpdateDate); err != nil {
			return nil, err
		} else {
			allHRContacts = append(allHRContacts, tempHRContact)
			log.Println("HRContact.id: ", &tempHRContact.Id)
		}
	}
	return allHRContacts, nil
}
