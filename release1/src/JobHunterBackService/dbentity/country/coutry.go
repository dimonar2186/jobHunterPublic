package country

//Basic structure description of country and basic methods are provided here.
// Countries is a dictionary. So I do not need to create/update/delete entites at this moment

import (
	entity "JobHunterBackService/dbentity"
	"database/sql"
	"fmt"
	"log"
)

const (
	countryTable   string = "public.countries"
	locationsTable string = "public.locations"
)

var tableTypes = map[string]string{
	"*vacancy.Vacancy":                         "vacancies",
	"*employer.Employer":                       "employers",
	"*jobsearchingProcess.JobSearchingProcess": "job_searching_processes",
	"*offer.Offer":                             "offers"}

type Country struct {
	Id   string
	Name string
	Code string
}

// Country constructor
func NewCountry(id, name, code string) *Country {
	return &Country{
		Id:   id,
		Name: name,
		Code: code}
}

// country select from DB record
func (country *Country) FindById(id string, tx *sql.Tx) error {
	query := fmt.Sprintf("SELECT id, name, code FROM %s WHERE id = $1", countryTable)
	rows, err := tx.Query(query, id) //read the country from the database
	if err != nil {
		log.Println("Failed to find country by id: ", id)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&country.Id, &country.Name, &country.Code); err != nil {
			return err
		} else {
			log.Println("Failed to extract country with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("country with id=%q not found", id)
}

// Saving country to DB is impossible
func (country *Country) Save(tx *sql.Tx) error {
	return fmt.Errorf("can not save country")
}

// Deleting the country from DB is impossible
func (country *Country) Delete(tx *sql.Tx) error {
	return fmt.Errorf("can not delete country")
}

func (country *Country) StatusUpdate(status string, tx *sql.Tx) error {
	//It is impossible to change country status now, because there is no such attribure
	return fmt.Errorf("can not update status of country")
}

func (country *Country) Update(tx *sql.Tx) error {
	//It is impossible to change coutry now, because there is no such need
	return fmt.Errorf("can not update country")
}

// Getter for Id
func (country *Country) GetId() string {
	return country.Id
}

// Get all countries
func GetCountries(tx *sql.Tx) (countries []*Country, err error) {
	query := fmt.Sprintf("SELECT id, name, code FROM %s", countryTable)
	rows, err := tx.Query(query) //read all countries from the database
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty country
		tempCountry := NewCountry("", "", "")
		//fulfill this temp country with row data and then append it into the slice of countries
		if err := rows.Scan(&tempCountry.Id, &tempCountry.Name, &tempCountry.Code); err != nil {
			log.Println("scan failed")
			return nil, err
		} else {
			countries = append(countries, tempCountry)
		}
	}
	return countries, nil
}

func GetAllLocations(object entity.Entity, tx *sql.Tx) (locations []*Country, err error) {
	tableType, ok := tableTypes[fmt.Sprintf("%T", object)]
	if ok {
		query := fmt.Sprintf("SELECT cont.id, cont.name, cont.code FROM %s loc join %s cont on loc.country = cont.id where loc.object_id = $1 and table_type = $2", locationsTable, countryTable)
		rows, err := tx.Query(query, object.GetId(), tableType) //read countries of the object from the database
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		tempCountry := NewCountry("", "", "")
		if rows.Next() {
			if err = rows.Scan(tempCountry.Id, tempCountry.Name, tempCountry.Code); err != nil {
				return nil, err
			} else {
				locations = append(locations, tempCountry)
			}
		} else {
			return nil, nil
		}
		return locations, nil
	} else {
		return nil, fmt.Errorf("Country selection. Unkown object type")
	}
}

func DeleteAllCountries(object entity.Entity, tx *sql.Tx) error {
	tableType, ok := tableTypes[fmt.Sprintf("%T", object)]
	if ok {
		query := "DELETE FROM " + locationsTable + " WHERE object_id = $1 AND table_type = $2"
		log.Printf("Executing query: %s with parameters: %v, %v", query, object.GetId(), tableType)
		_, err := tx.Exec(query, object.GetId(), tableType) //delete locations of the object from the database
		//_, err := tx.Exec("DELETE FROM public.locations WHERE object_id = '0c34423b-39df-43b0-9527-23ae5f92be8e' AND table_type = 'employers'")
		if err != nil {
			log.Println("location deletion failed ", err)
			return err
		}
		return nil
	} else {
		return fmt.Errorf("locations deletion. unkown object type")
	}
}

func (country Country) SaveLocation(object entity.Entity, tx *sql.Tx) error {
	tableType, ok := tableTypes[fmt.Sprintf("%T", object)]
	if ok {
		if checked, err := CheckCountryId(country.Id, tx); !checked || err != nil {
			return fmt.Errorf("unkown country")
		}
		query := fmt.Sprintf("INSERT INTO %s (object_id, table_type, country) VALUES ($1, $2, $3) RETURNING id", locationsTable)
		err := tx.QueryRow(query, object.GetId(), tableType, country.Id).Scan(&country.Id) //write countries of the object to the database and get an id
		if err != nil {
			return fmt.Errorf("country saving. Saving failed %w", err)
		}
		return nil
	} else {
		return fmt.Errorf("country saving. unkown object type")
	}
}

// Check existance of country checks if there is an id in the country table
func CheckCountryId(id string, tx *sql.Tx) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id= $1)", countryTable)
	log.Println(query)
	var checked bool
	rows, err := tx.Query(query, id)
	if err != nil {
		log.Println("Failed to check country with id: ", id)
		return false, fmt.Errorf("failed to check country with id: %s %w", id, err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&checked)
		if checked {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, fmt.Errorf("failed to check country with id: %s", id)
}
