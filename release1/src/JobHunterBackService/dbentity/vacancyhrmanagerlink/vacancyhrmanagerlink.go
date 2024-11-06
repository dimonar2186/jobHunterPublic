package vacancyhrmanagerlink

import (
	entity "JobHunterBackService/dbentity"
	"database/sql"
	"fmt"
)

const (
	vacanciesHrmanagersLinks = "public.vacancies_to_hrmanagers"
)

func Save(vacancy entity.Entity, hrmanager entity.Entity, tx *sql.Tx) (id string, err error) {
	query := fmt.Sprintf("INSERT INTO %s (hrmanager, vacancy) VALUES $1, $2 RETURNING id",
		vacanciesHrmanagersLinks)

	//Define vacancy as vacancy.Vacancy
	if fmt.Sprintf("%T", vacancy) != "*vacancy.Vacancy" {
		return "", fmt.Errorf("wrong object provided as vacancy")
	}

	//Define hrmanager as hrmanager.HRManager
	if fmt.Sprintf("%T", vacancy) != "*hrmanager.HRManager" {
		return "", fmt.Errorf("wrong object provided as hrmanager")
	}

	//Insert into DB
	if err := tx.QueryRow(query, vacancy.GetId(), hrmanager.GetId()).Scan(&id); err != nil {
		return "", fmt.Errorf("failed to save link between vacacy.id %s and hrmanager.id %s %w", vacancy.GetId(), hrmanager.GetId(), err)
	}
	return id, nil
}

func Delete(object entity.Entity, tx *sql.Tx) error {
	var query string = "DELETE FROM " + vacanciesHrmanagersLinks
	objectType := fmt.Sprintf("%T", object)
	switch objectType {
	case "*vacancy.Vacancy":
		query = query + " where vacancy = $1"
	case "*hrmanager.HRManager":
		query = query + " where hrmanager = $1"
	}
	if _, err := tx.Exec(query, object.GetId()); err != nil { //delete links of the object from the database
		return err
	}
	return nil
}

func GetByObjectId(object entity.Entity, tx *sql.Tx) (ids []string, err error) {
	type selectParams struct {
		idField, searchField string
	}
	searchMatrix := make(map[string]selectParams)
	searchMatrix["*hrmanager.HRManager"] = selectParams{idField: "vacnacy", searchField: "hrmanager"}
	searchMatrix["*vacancy.Vacancy"] = selectParams{idField: "hrmanager", searchField: "vacnacy"}

	rows, err := tx.Query("SELECT $1 FROM $2 where $3 = $4", searchMatrix[fmt.Sprintf("%T", object)].idField, vacanciesHrmanagersLinks, searchMatrix[fmt.Sprintf("%T", object)].searchField, object.GetId()) //read list from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		} else {
			ids = append(ids, id)
		}
	}
	return ids, nil
}
