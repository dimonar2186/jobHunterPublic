package applicationstage

//Basic structure description of applicationStage and basic methods are provided here.
//Important: Save, Update, Delete, StatusUpdate, GetAllNeighbourApplicationStages, GetAllApplicationStagesOfVacancy require an opened transaction to be executed

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	applicationStageTable string = "public.application_stages"
)

type ApplicationStage struct {
	Id           string
	VacancyId    string
	IndexNumber  int
	Name         string
	Status       string
	CreationDate time.Time
	UpdateDate   time.Time
}

// applicationStage constructor
func NewApplicationStage(
	id, vacancyId, name, status string,
	indexNumber int,
	creationDate, updateDate time.Time) *ApplicationStage {
	return &ApplicationStage{
		Id:           id,
		VacancyId:    vacancyId,
		IndexNumber:  indexNumber,
		Name:         name,
		Status:       status,
		CreationDate: creationDate,
		UpdateDate:   updateDate}
}

// applicationStage select from DB record
func (applicationStage *ApplicationStage) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, vacancy, index_number, name, status, created_at, updated_at FROM $1 where id = $2", applicationStageTable, id) //read the vacancy from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(applicationStage.Id, applicationStage.VacancyId, applicationStage.IndexNumber, applicationStage.Name,
			applicationStage.Status, applicationStage.CreationDate, applicationStage.UpdateDate); err != nil {
			return err
		} else {
			log.Println("Failed to extract applicationStage with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("applicationStage with id=%q not found", id)
}

// Saving applicationStage to DB
func (applicationStage *ApplicationStage) Save(tx *sql.Tx) error {
	//form the query
	applicationStageDBCols := "(vacancy, index_number, name, status, created_at, updated_at)"
	applicationStageInsertParameters := "($1, $2, $3, $4, $5, $6)"
	query := fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s RETURNING id",
		applicationStageTable, applicationStageDBCols, applicationStageInsertParameters,
	)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query,
		applicationStage.VacancyId, applicationStage.IndexNumber, applicationStage.Name,
		applicationStage.Status, applicationStage.CreationDate, time.Now()).Scan(&id); err != nil {
		return fmt.Errorf("failed to save applicationStage %w", err)
	}
	applicationStage.Id = id
	return nil
}

// Deleting the applicationStage from DB
func (applicationStage *ApplicationStage) Delete(tx *sql.Tx) error {
	//form deleting query and execute it
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = %s", applicationStageTable, applicationStage.Id)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("failed to delete the applicationStage %w", err)
	}
	return nil
}

func (applicationStage *ApplicationStage) StatusUpdate(status string, tx *sql.Tx) error {
	//Form and execute status updating query
	query := fmt.Sprintf("UPDATE %s SET status = %s, updated_at = %s WHERE id = %s", applicationStageTable, status, time.Now(), applicationStage.Id)
	if _, err := tx.Exec(query); err != nil {
		errorMessage := "failed to update status of the applicationStage with id=" + applicationStage.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

func (applicationStage *ApplicationStage) Update(tx *sql.Tx) error {
	//Form and execute update query
	applicationStageDBCols := "(vacancy, index_number, name, status, created_at, updated_at)"
	applicationStageInsertParameters := "($1, $2, $3, $4, $5, $6)"
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE id = %s", applicationStageTable, applicationStageDBCols, applicationStageInsertParameters, applicationStage.Id)
	if _, err := tx.Exec(query,
		applicationStage.VacancyId, applicationStage.IndexNumber, applicationStage.Name,
		applicationStage.Status, applicationStage.CreationDate, time.Now()); err != nil {
		errorMessage := "failed to update the applicationStage with id=" + applicationStage.Id + "%w"
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Getter for Id
func (applicationStage *ApplicationStage) GetId() string {
	return applicationStage.Id
}

// form full list of ApplicationStage related to the specific one
func (applicationStage *ApplicationStage) GetAllNeighbourApplicationStages(tx *sql.Tx) (allApplicationStages []*ApplicationStage, err error) {
	//Form a slice of applicationStages, those are connected with the used one

	rows, err := tx.Query("SELECT id, vacancy, index_number, name, status, created_at, updated_at FROM $1 where vacancy = $2", applicationStageTable, applicationStage.VacancyId) //read the applicationStages list from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempId, tempVacancyId, tempName, tempStatus string
		var tempIndexNumber int
		var tempCreationDate, tempUpdateDate time.Time
		if err := rows.Scan(tempId, tempVacancyId, tempIndexNumber, tempName, tempStatus, tempCreationDate, tempUpdateDate); err != nil {
			return nil, err
		} else {
			tempApplicationStage := NewApplicationStage(tempId, tempVacancyId, tempName, tempStatus, tempIndexNumber, tempCreationDate, tempUpdateDate)
			allApplicationStages = append(allApplicationStages, tempApplicationStage)
			log.Println("applicationStage.id: ", &applicationStage.Id)
		}
	}
	if len(allApplicationStages) == 0 {
		return nil, nil
	}
	return allApplicationStages, nil
}

// form full list of ApplicationStage for the vacancyId
func GetAllApplicationStagesOfVacancy(vacancyId string, tx *sql.Tx) (allApplicationStages []*ApplicationStage, err error) {

	rows, err := tx.Query("SELECT id, vacancy, index_number, name, status, created_at, updated_at FROM $1 where vacancy = $2", applicationStageTable, vacancyId) //read the applicationStages list from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty temporary applicationStage
		tempApplicationStage := NewApplicationStage("", "", "", "", 0, time.Time{}, time.Time{})
		//fulfill this temp applicationStage with row data and then append it into the slice of applicationStages
		if err := rows.Scan(tempApplicationStage.Id, tempApplicationStage.VacancyId, tempApplicationStage.IndexNumber,
			tempApplicationStage.Name, tempApplicationStage.Status, tempApplicationStage.CreationDate,
			tempApplicationStage.UpdateDate); err != nil {
			return nil, err
		} else {
			allApplicationStages = append(allApplicationStages, tempApplicationStage)
			log.Println("applicationStage.id: ", &tempApplicationStage.Id)
		}
	}
	return allApplicationStages, nil
}
