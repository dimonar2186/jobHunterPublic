package user

//Basic structure description of user and basic methods are provided here.
//Important: GetAllVacancies, GetAllJobSearchingProcesses require an opened transaction to be executed

import (
	"JobHunterBackService/dbentity/jobsearchingprocess"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

const (
	userTable string = "public.users"
)

type User struct {
	Id string
}

// user constructor
func NewUser(id string) *User {
	return &User{
		Id: id}
}

// user select from DB record
func (user *User) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id FROM $1 where id = $2", userTable, id) //read the user from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(user.Id); err != nil {
			return err
		} else {
			log.Println("Failed to extract user with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("user with id=%q not found", id)
}

// Saving user to DB
func (user *User) Save(tx *sql.Tx) error {
	//form the query
	userDBCols := "(id)"
	userInsertParameters := "($1)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		userTable, userDBCols, userInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query, user.Id).Scan(&id); err != nil {
		return fmt.Errorf("failed to save user %w", err)
	}
	user.Id = id
	return nil
}

// Deleting the user from DB
func (user *User) Delete(tx *sql.Tx) error {
	//form deleting query and execute it
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = %s", userTable, user.Id)
	if _, err := tx.Query(query); err != nil {
		return fmt.Errorf("failed to delete the user %w", err)
	}
	return nil
}

func (user *User) StatusUpdate(status string, tx *sql.Tx) error {
	//It is impossible to change user status now, because there is no such attribure
	return nil
}

func (user *User) Update(tx *sql.Tx) error {
	//It is impossible to change user now, because there is no such need
	return nil
}

// Get all jobSearchingProcesses created by user
func (user *User) GetUserJobSearchingProcesses(tx *sql.Tx) (userJobSearchingProcesses []*jobsearchingprocess.Jobsearchingprocess, err error) {
	rows, err := tx.Query("SELECT  id, name, min_mounthly_salary, max_mounthly_salary, \"position\", \"user\", created_at, updated_at, status, is_deleted, deleted_at, currency FROM $1 where user = $2", "public.job_searching_processes", user.Id) //read all jobSearchingProcesses from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty jobSearchingProcess
		tempJobSearchingProcess := jobsearchingprocess.NewJobSearchingProcess("", "", "", "", "", 0, 0, time.Time{}, time.Time{}, time.Time{}, false, []string{})
		//fulfill this temp jobSearchingProcess with row data and then append it into the slice of jobSearchingProcesses
		var dbPositions pq.StringArray
		if err := rows.Scan(tempJobSearchingProcess.Id, tempJobSearchingProcess.Name, tempJobSearchingProcess.MinimumMonthlySalary, tempJobSearchingProcess.MaximumMonthlySalary, &dbPositions, tempJobSearchingProcess.User, tempJobSearchingProcess.CreationDate, tempJobSearchingProcess.UpdateDate, tempJobSearchingProcess.Status, tempJobSearchingProcess.IsDeleted, tempJobSearchingProcess.Currency); err != nil {
			return nil, err
		} else {
			tempJobSearchingProcess.Positions = dbPositions
			userJobSearchingProcesses = append(userJobSearchingProcesses, tempJobSearchingProcess)
			log.Println("jobSearchingProcess.id: ", &tempJobSearchingProcess.Id)
		}
	}
	return userJobSearchingProcesses, nil
}
