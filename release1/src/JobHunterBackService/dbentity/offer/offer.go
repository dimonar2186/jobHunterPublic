package offer

import (
	"JobHunterBackService/dbentity/country"
	"JobHunterBackService/dbentity/jobtype"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	offerTable string = "public.offers"
)

type Offer struct {
	Id             string
	VacancyId      string
	MonthlySalary  float64
	Currency       string
	ContractLength int
	Position       string
	Commentary     string
	CreationDate   time.Time
	UpdateDate     time.Time
	JobType        string
}

// jobSearchingProcess constructor
func NewOffer(
	id, vacancyId, currency, position, commentary, jobType string,
	monthlySalary float64,
	creationDate, updateDate time.Time,
	contractLength int) *Offer {
	return &Offer{
		Id:             id,
		VacancyId:      vacancyId,
		MonthlySalary:  monthlySalary,
		Currency:       currency,
		ContractLength: contractLength,
		Position:       position,
		Commentary:     commentary,
		CreationDate:   creationDate,
		UpdateDate:     updateDate,
		JobType:        jobType}
}

// jobSearchingProcess select from DB record
func (offer *Offer) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, vacancy, mounthly_salary, currency, contract_length, \"position\", commentary, created_at, updated_at, job_type FROM $1 where id = $2", offerTable, id) //read the offer from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(offer.Id, offer.VacancyId, offer.MonthlySalary, offer.Currency, offer.ContractLength, offer.Position, offer.Commentary, offer.CreationDate, offer.UpdateDate, offer.JobType); err != nil {
			return err
		} else {
			log.Println("Failed to extract offer with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("Offer with id=%q not found", id)
}

// Saving jobSearchingProcess to DB
func (offer *Offer) Save(tx *sql.Tx) error {
	//form the query
	offerDBCols := "(vacancy, mounthly_salary, currency, contract_length, \"position\", commentary, created_at, updated_at, job_type)"
	offerInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		offerTable, offerDBCols, offerInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		offer.VacancyId, offer.MonthlySalary, offer.Currency, offer.ContractLength, offer.Position, offer.Commentary,
		time.Now(), time.Now(), offer.JobType).Scan(&id); err != nil {
		return fmt.Errorf("failed to save offer %w", err)
	}
	offer.Id = id
	return nil
}

func (offer *Offer) Delete(tx *sql.Tx) error {
	//delete all links to jobTypes
	jobtype.DeleteAllJobTypes(offer, tx)

	//delete all locations
	country.DeleteAllCountries(offer, tx)

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = %s",
		offerTable, offer.Id)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("failed to delete offer with id=%s %w", offer.Id, err)
	}
	return nil
}

func (offer *Offer) StatusUpdate(status string, tx *sql.Tx) error {
	//Nothing to do with offer's status, because there is no status atribute
	return nil
}

func (offer *Offer) Update(tx *sql.Tx) error {
	//Form and execute update query
	offerDBCols := "(vacancy, mounthly_salary, currency, contract_length, \"position\", commentary, updated_at, job_type)"
	offerInsertParameters := "($1, $2, $3, $4, $5, $6, $7, $8)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", offerTable, offerDBCols, offerInsertParameters, offer.Id)
	if _, err := tx.Exec(query, offer.VacancyId, offer.MonthlySalary, offer.Currency, offer.ContractLength, offer.Position, offer.Commentary, time.Now(), offer.JobType); err != nil {
		return fmt.Errorf("failed to update offer with id=%s %w", offer.Id, err)
	}
	return nil
}

// Getter for Id
func (offer *Offer) GetId() string {
	return offer.Id
}
