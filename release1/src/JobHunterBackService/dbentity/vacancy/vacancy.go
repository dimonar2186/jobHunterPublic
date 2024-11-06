package vacancy

import (
	"JobHunterBackService/dbentity/applicationstage"
	"JobHunterBackService/dbentity/country"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	vacancyTable   string = "public.vacancies"
	tableType      string = "vacancies"
	locationsTable string = "public.locations"
)

type Vacancy struct {
	Id                    string
	JobSearchingProcessId string
	Currency              string
	ContractLength        float64
	Benefits              string
	Responsibilities      string
	Position              string
	Commentary            string
	EmployerId            string
	ApllicationDate       time.Time
	CreationDate          time.Time
	UpdateDate            time.Time
	Status                string
	MinimumMonthlySalary  float64
	MaximumMonthlySalary  float64
}

// vacancy constructor
func NewVacancy(
	id, jobSearchingProcessId, currency, benefits, responsibilities, position, commentary, employerId, status string,
	contractLength, minimumMonthlySalary, maximumMonthlySalary float64,
	apllicationDate, creationDate, updateDate time.Time) *Vacancy {
	return &Vacancy{
		Id:                    id,
		JobSearchingProcessId: jobSearchingProcessId,
		Currency:              currency,
		ContractLength:        contractLength,
		Benefits:              benefits,
		Responsibilities:      responsibilities,
		Position:              position,
		Commentary:            commentary,
		EmployerId:            employerId,
		ApllicationDate:       apllicationDate,
		CreationDate:          creationDate,
		UpdateDate:            updateDate,
		Status:                status,
		MinimumMonthlySalary:  minimumMonthlySalary,
		MaximumMonthlySalary:  maximumMonthlySalary}
}

// Vacancy select from DB record
func (vacancy *Vacancy) FindById(id string, tx *sql.Tx) error {

	rows, err := tx.Query("SELECT id, job_searching_process, currency, contract_length, benefits, responsibilities, \"position\", commentary, employer, applied_at, created_at, updated_at, status, min_mounthly_salary, max_mounthly_salary FROM $1 where id = $2", vacancyTable, id) //read the vacancy from the database
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(vacancy.Id, vacancy.JobSearchingProcessId, vacancy.Currency, vacancy.ContractLength, vacancy.Benefits,
			vacancy.Responsibilities, vacancy.Position, vacancy.Commentary, vacancy.EmployerId, vacancy.ApllicationDate,
			vacancy.CreationDate, vacancy.UpdateDate, vacancy.Status, vacancy.MinimumMonthlySalary, vacancy.MaximumMonthlySalary); err != nil {
			return err
		} else {
			log.Println("Failed to extract vacancy with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("Vacancy with id=%q not found", id)
}

// Saving vacancy to DB
func (vacancy *Vacancy) Save(tx *sql.Tx) error {

	//form the query
	vacancyDBCols := "(job_searching_process, currency, contract_length, benefits, responsibilities, \"position\", commentary, employer, applied_at, created_at, updated_at, status, min_mounthly_salary, max_mounthly_salary)"
	vacancyInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		vacancyTable, vacancyDBCols, vacancyInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		vacancy.JobSearchingProcessId, vacancy.Currency, vacancy.ContractLength, vacancy.Benefits,
		vacancy.Responsibilities, vacancy.Position, vacancy.Commentary, vacancy.EmployerId, vacancy.ApllicationDate,
		vacancy.CreationDate, time.Now(), vacancy.Status, vacancy.MinimumMonthlySalary,
		vacancy.MaximumMonthlySalary).Scan(&id); err != nil {
		return fmt.Errorf("failed to save vacancy %w", err)
	}
	vacancy.Id = id
	return nil
}

func (vacancy *Vacancy) Delete(tx *sql.Tx) error {
	//soft vacancy deletion
	vacancy.Status = "deleted"
	if err := vacancy.Update(tx); err != nil {
		return fmt.Errorf("failed to delete vacancy %w", err)
	}

	//Hard delete was commented because vacancy must be deleted softly
	/*
		//delete all links to jobTypes
		if err := jobtype.DeleteAllJobTypes(vacancy, tx); err != nil {
			return err
		}

		//delete all locations
		if err := country.DeleteAllCountries(vacancy, tx); err != nil {
			return err
		}

		//delete links with hrmanagers
		if err := vacancyhrmanagerlink.Delete(vacancy, tx); err != nil {
			return err
		}

		//form deleting query and execute it
		query := fmt.Sprintf(
			"DELETE FROM %s WHERE id = %s",
			vacancyTable, vacancy.Id)
		if _, err := tx.Query(query); err != nil {
			return fmt.Errorf("failed to delete vacancy %w", err)
		}
	*/
	return nil
}

func (vacancy *Vacancy) StatusUpdate(status string, tx *sql.Tx) error {

	//Form and execute status updating query
	query := fmt.Sprintf("UPDATE %s SET status = %s, updated_at = %s WHERE id = %s", vacancyTable, status, time.Now(), vacancy.Id)
	if _, err := tx.Exec(query); err != nil {
		errorMessage := "failed to update status of the vacancy with id=" + vacancy.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

func (vacancy *Vacancy) Update(tx *sql.Tx) error {

	//Form and execute update query
	vacancyDBCols := "(job_searching_process, currency, contract_length, benefits, responsibilities, \"position\", commentary, employer, applied_at, created_at, updated_at, status, min_mounthly_salary, max_mounthly_salary)"
	vacancyInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", vacancyTable, vacancyDBCols, vacancyInsertParameters, vacancy.Id)
	if _, err := tx.Exec(query,
		vacancy.JobSearchingProcessId, vacancy.Currency, vacancy.ContractLength, vacancy.Benefits,
		vacancy.Responsibilities, vacancy.Position, vacancy.Commentary, vacancy.EmployerId, vacancy.ApllicationDate,
		vacancy.CreationDate, time.Now(), vacancy.Status, vacancy.MinimumMonthlySalary,
		vacancy.MaximumMonthlySalary); err != nil {
		errorMessage := "failed to update the vacancy with id=" + vacancy.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Getter for Id
func (vacancy *Vacancy) GetId() string {
	return vacancy.Id
}

// Create default set of applicationStages for the vacancy
func (vacancy *Vacancy) CreateDefaultApplicationStages(tx *sql.Tx) error {
	//create slice with default set of applicationStage names
	defaultNames := map[int]string{
		0: "Applied",
		1: "HRContact",
		2: "Interview",
		3: "Feedback",
	}

	for tempIndex := 0; tempIndex < len(defaultNames); tempIndex++ {
		tempApplicationStage := applicationstage.NewApplicationStage("", vacancy.Id, defaultNames[tempIndex], "", tempIndex, time.Now(), time.Now())
		if err := tempApplicationStage.AddApplicationStage(); err != nil {
			return err
		}
	}
	return nil
}

// Collect all locations of the vacancy
func (vacancy *Vacancy) GetAllVacancyLocations(tx *sql.Tx) (locations []*country.Country, err error) {

	rows, err := tx.Query("SELECT SELECT cont.id, cont.name, cont.code FROM public.locations loc join public.countries cont on loc.country = cont.id where loc.object_id = $2 and table_type = $3", locationsTable, vacancy.Id, tableType) //read countries of the vacancy from the database
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tempCountry := country.NewCountry("", "", "")
	if rows.Next() {
		if err = rows.Scan(tempCountry.Id, tempCountry.Name, tempCountry.Code); err != nil {
			return nil, err
		} else {
			locations = append(locations, tempCountry)
		}
	}
	return locations, nil
}
