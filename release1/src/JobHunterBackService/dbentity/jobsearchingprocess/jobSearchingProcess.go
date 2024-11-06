package jobsearchingprocess

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	"JobHunterBackService/dbentity/vacancy"

	"github.com/lib/pq"
)

const (
	jobSearchingProcessTable string = "public.job_searching_processes"
)

type Jobsearchingprocess struct {
	Id                   string
	Name                 string
	MinimumMonthlySalary float64
	MaximumMonthlySalary float64
	Positions            []string
	User                 string
	CreationDate         time.Time
	UpdateDate           time.Time
	Status               string
	IsDeleted            bool
	DeleteDate           time.Time
	Currency             string
}

// jobSearchingProcess constructor
func NewJobSearchingProcess(
	id, name, user, status, currency string,
	minimumMonthlySalary, maximumMonthlySalary float64,
	creationDate, updateDate, deleteDate time.Time,
	isDeleted bool,
	positions []string) *Jobsearchingprocess {
	return &Jobsearchingprocess{
		Id:                   id,
		Name:                 name,
		MinimumMonthlySalary: minimumMonthlySalary,
		MaximumMonthlySalary: maximumMonthlySalary,
		Positions:            positions,
		User:                 user,
		CreationDate:         creationDate,
		UpdateDate:           updateDate,
		IsDeleted:            isDeleted,
		DeleteDate:           deleteDate,
		Status:               status,
		Currency:             currency}
}

// jobSearchingProcess select from DB record
func (jobSearchingProcess *Jobsearchingprocess) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, name, min_mounthly_salary, max_mounthly_salary, \"position\", \"user\", created_at, updated_at, status, is_deleted, deleted_at, currency FROM $1 where id = $2", jobSearchingProcessTable, id) //read the jobSearchingProcess from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(jobSearchingProcess.Id, jobSearchingProcess.Name,
			jobSearchingProcess.MinimumMonthlySalary, jobSearchingProcess.MaximumMonthlySalary,
			jobSearchingProcess.Positions, jobSearchingProcess.User, jobSearchingProcess.CreationDate,
			jobSearchingProcess.UpdateDate, jobSearchingProcess.Status, jobSearchingProcess.IsDeleted,
			jobSearchingProcess.DeleteDate, jobSearchingProcess.Currency); err != nil {
			return err
		} else {
			log.Println("Failed to extract jobSearchingProccess with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("jobSearchingProcess with id=%q not found", id)
}

// Saving jobSearchingProcess to DB
func (jobSearchingProcess *Jobsearchingprocess) Save(tx *sql.Tx) error {
	//form the query
	jobSearchingProcessDBCols := "(name, min_mounthly_salary, max_mounthly_salary, \"position\", \"user\", created_at, updated_at, is_deleted, deleted_at, currency, status)"
	jobSearchingInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		jobSearchingProcessTable, jobSearchingProcessDBCols, jobSearchingInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		jobSearchingProcess.Name, jobSearchingProcess.MinimumMonthlySalary, jobSearchingProcess.MaximumMonthlySalary,
		pq.Array(jobSearchingProcess.Positions), jobSearchingProcess.User, jobSearchingProcess.CreationDate,
		time.Now(), jobSearchingProcess.IsDeleted, jobSearchingProcess.DeleteDate,
		jobSearchingProcess.Currency, jobSearchingProcess.Status).Scan(&id); err != nil {
		return fmt.Errorf("failed to save jobSearchingProcess %w", err)
	}
	jobSearchingProcess.Id = id
	return nil
}

func (jobSearchingProcess *Jobsearchingprocess) Delete(tx *sql.Tx) error {
	jobSearchingProcess.DeleteDate = time.Now()
	jobSearchingProcess.IsDeleted = true
	jobSearchingProcess.Status = "deleted"

	if err := jobSearchingProcess.Update(tx); err != nil {
		return fmt.Errorf("failed to delete jobSearchingProcess %w", err)
	}

	return nil
}

func (jobSearchingProcess *Jobsearchingprocess) StatusUpdate(status string, tx *sql.Tx) error {
	//Form and execute status updating query
	query := fmt.Sprintf("UPDATE %s SET status = %s, updated_at = %s WHERE id = %s", jobSearchingProcessTable, status, time.Now(), jobSearchingProcess.Id)
	if _, err := tx.Exec(query); err != nil {
		errorMessage := "failed to update status of the jobSearchingProcess with id=" + jobSearchingProcess.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

func (jobSearchingProcess *Jobsearchingprocess) Update(tx *sql.Tx) error {
	//Form and execute update query
	jobSearchingProcessDBCols := "(name, min_mounthly_salary, max_mounthly_salary, \"position\", \"user\", created_at, updated_at, is_deleted, deleted_at, currency, status)"
	jobSearchingInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", jobSearchingProcessTable, jobSearchingProcessDBCols, jobSearchingInsertParameters, jobSearchingProcess.Id)
	if _, err := tx.Exec(query,
		jobSearchingProcess.Name, jobSearchingProcess.MinimumMonthlySalary, jobSearchingProcess.MaximumMonthlySalary,
		pq.Array(jobSearchingProcess.Positions), jobSearchingProcess.User, jobSearchingProcess.CreationDate,
		time.Now(), jobSearchingProcess.IsDeleted, jobSearchingProcess.DeleteDate,
		jobSearchingProcess.Currency, jobSearchingProcess.Status); err != nil {
		return fmt.Errorf("failed to update jobSearchingProcess %w", err)
	}
	return nil
}

// Getter for Id
func (jobSearchingProcess *Jobsearchingprocess) GetId() string {
	return jobSearchingProcess.Id
}

// form full list of vacancies for the jobSearchingProcess
func (jobSearchingProcess *Jobsearchingprocess) GetAllVscsnciesOfJobSearchingProcess(tx *sql.Tx) (allVacancies []*vacancy.Vacancy, err error) {
	vacancyDBCols := "(id, job_searching_process, currency, contract_length, benefits, responsibilities, \"position\", commentary, employer, applied_at, created_at, updated_at, status, min_mounthly_salary, max_mounthly_salary)"
	vacancyTable := "public.vacancies"
	rows, err := tx.Query("SELECT $1 FROM $2 where job_searching_process = $3", vacancyDBCols, vacancyTable, jobSearchingProcess.Id) //read the vacancies list from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty temporary HRContact
		tempVacancy := vacancy.NewVacancy("", "", "", "", "", "", "", "", "", float64(0), float64(0), float64(0), time.Time{}, time.Time{}, time.Time{})
		//fulfill this temp Vacancies with row data and then append it into the slice of Vacancies
		if err := rows.Scan(tempVacancy.Id, tempVacancy.JobSearchingProcessId, tempVacancy.Currency, tempVacancy.ContractLength, tempVacancy.Benefits,
			tempVacancy.Responsibilities, tempVacancy.Position, tempVacancy.Commentary, tempVacancy.EmployerId, tempVacancy.ApllicationDate,
			tempVacancy.CreationDate, tempVacancy.UpdateDate, tempVacancy.Status, tempVacancy.MinimumMonthlySalary, tempVacancy.MaximumMonthlySalary); err != nil {
			return nil, err
		} else {
			allVacancies = append(allVacancies, tempVacancy)
			log.Println("HRContact.id: ", &tempVacancy.Id)
		}
	}
	return allVacancies, nil
}
