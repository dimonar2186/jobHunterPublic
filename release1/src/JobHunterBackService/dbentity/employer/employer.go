package employer

//Basic structure description of employer and basic methods are provided here.
//Important: Save, Update, Delete, StatusUpdate, GetAllHrcontactsOfHrmanager require an opened transaction to be executed

import (
	"JobHunterBackService/dbentity/country"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	employersTable string = "public.employers"
)

type Employer struct {
	Id           string
	Name         string
	CreationDate time.Time
}

// employer constructor
func NewEmployer(
	id, name string,
	creationDate time.Time) *Employer {
	return &Employer{
		Id:           id,
		Name:         name,
		CreationDate: creationDate}
}

// employer select from DB record
func (employer *Employer) FindById(id string, tx *sql.Tx) error {
	query := "SELECT id, name, created_at FROM " + employersTable + " where id = $1"
	rows, err := tx.Query(query, id) //read the employer from the database

	if err != nil {
		log.Printf("employer with id=%s not found", id)
		return fmt.Errorf("employer with id=%q not found", id)
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&employer.Id, &employer.Name, &employer.CreationDate); err != nil {
			log.Println("Failed to extract employer with id: ", id)
			return err
		}

	}
	return nil
}

// Saving employer to DB
func (employer *Employer) Save(tx *sql.Tx) error {
	//form the query
	employerDBCols := "(name, created_at)"
	employerParameters := "($1, $2)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		employersTable, employerDBCols, employerParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query, &employer.Name, time.Now()).Scan(&id); err != nil {
		return fmt.Errorf("failed to save employer %w", err)
	}
	employer.CreationDate = time.Now()
	employer.Id = id
	return nil
}

// Deleting the employer from DB
func (employer *Employer) Delete(tx *sql.Tx) error {
	//delete all locations
	country.DeleteAllCountries(employer, tx)

	//form deleting query and execute it
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", employersTable)
	if _, err := tx.Exec(query, employer.Id); err != nil {
		return fmt.Errorf("failed to delete the employer %w", err)
	}
	return nil
}

func (employer *Employer) StatusUpdate(status string, tx *sql.Tx) error {
	//There is no field called status.
	return fmt.Errorf("can't update status of the employer")
}

func (employer *Employer) Update(tx *sql.Tx) error {
	//Form and execute update query
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = '%s'", employersTable, employer.Id)
	if _, err := tx.Exec(query, employer.Name); err != nil {
		errorMessage := fmt.Sprintf("failed to update the employer with id="+employer.Id+"%w", err)
		return fmt.Errorf(errorMessage, err)
	}

	return nil
}

// Getter for Id
func (employer *Employer) GetId() string {
	return employer.Id
}
