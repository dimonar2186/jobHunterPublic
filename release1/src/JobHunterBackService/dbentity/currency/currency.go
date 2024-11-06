package currency

//Basic structure description of currency and basic methods are provided here.
// Currencies is a dictionary. So I do not need to create/update/delete entites at this moment

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	currencyTable string = "public.currencies"
)

type Currency struct {
	Id        string
	Name      string
	ShortName string
	IsoCode   string
}

// Currency constructor
func NewCurrency(id, name, shortName, isoCode string) *Currency {
	return &Currency{
		Id:        id,
		Name:      name,
		ShortName: shortName,
		IsoCode:   isoCode}
}

// currency select from DB record
func (currency *Currency) FindById(id string, tx *sql.Tx) error {
	query := fmt.Sprintf("SELECT id, name, short_name, iso_code FROM %s where id = $1", currencyTable)
	rows, err := tx.Query(query, id) //read the currency from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&currency.Id, &currency.Name, &currency.ShortName, &currency.IsoCode); err != nil {
			log.Println("Failed to extract currency with id: ", id)
			return err
		}
		return nil
	}
	return fmt.Errorf("currency with id=%q not found", id)
}

// Saving currency to DB is impossible
func (currency *Currency) Save(tx *sql.Tx) error {
	return nil
}

// Deleting the currency from DB is impossible
func (currency *Currency) Delete(tx *sql.Tx) error {
	return nil
}

func (currency *Currency) StatusUpdate(status string, tx *sql.Tx) error {
	//It is impossible to change currency status now, because there is no such attribure
	return nil
}

func (currency *Currency) Update(tx *sql.Tx) error {
	//It is impossible to change coutry now, because there is no such need
	return nil
}

// Get all currencies
func GetCurrencies(tx *sql.Tx) (currencies []*Currency, err error) {
	query := fmt.Sprintf("SELECT id, name, short_name, iso_code FROM %s", currencyTable)
	rows, err := tx.Query(query) //read all currencies from the database

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//Create an empty currency
		tempCurrency := NewCurrency("", "", "", "")
		//fulfill this temp currency with row data and then append it into the slice of currencies
		if err := rows.Scan(&tempCurrency.Id, &tempCurrency.Name, &tempCurrency.ShortName, &tempCurrency.IsoCode); err != nil {
			return nil, err
		} else {
			currencies = append(currencies, tempCurrency)
		}
	}
	return currencies, nil
}
