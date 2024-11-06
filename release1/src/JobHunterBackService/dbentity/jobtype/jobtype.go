package jobtype

// Collector of jobTypes for different objects

import (
	entity "JobHunterBackService/dbentity"
	"database/sql"
	"fmt"
)

const (
	linkToJobTypesTable string = "public.links_to_job_types"
)

var tableTypes = map[string]string{
	"*vacancy.Vacancy":                         "vacancies",
	"*employer.Employer":                       "employers",
	"*jobsearchingProcess.JobSearchingProcess": "job_searching_processes",
	"*offer.Offer":                             "offers"}

func GetAllJobTypes(object entity.Entity, tx *sql.Tx) (jobTypes []string, err error) {
	tableType, ok := tableTypes[fmt.Sprintf("%T", object)]
	if ok {
		rows, err := tx.Query("SELECT job_type FROM $1 where object_id = $2 and table_type = $3", linkToJobTypesTable, object.GetId(), tableType) //read jobTypes of the object from the database
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		//jobTypes = make([]string, 0)
		var tempJobType string
		if rows.Next() {
			if err = rows.Scan(&tempJobType); err != nil {
				return nil, err
			} else {
				jobTypes = append(jobTypes, tempJobType)
			}
		} else {
			return nil, nil
		}
		return jobTypes, nil
	} else {
		return nil, fmt.Errorf("country selection. unkown object type")
	}
}

func DeleteAllJobTypes(object entity.Entity, tx *sql.Tx) error {
	tableType, ok := tableTypes[fmt.Sprintf("%T", object)]
	if ok {
		_, err := tx.Query("DELETE FROM $1 where object_id = $2 and table_type = $3", linkToJobTypesTable, object.GetId(), tableType) //delete links to jobTypes of the object from the database
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("Links to jobTypes deletion. Unkown object type")
	}
}
