package messenger

//Basic structure description of hrcontact and basic methods are provided here.
//Important: Save, Update, Delete, StatusUpdate, GetAllHrcontactsOfHrmanager require an opened transaction to be executed

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	messengersTable string = "public.messengers"
)

type Messenger struct {
	Id   string
	Name string
}

// Messenger constructor
func NewMessenger(
	id, name string) *Messenger {
	return &Messenger{
		Id:   id,
		Name: name}
}

// Messenger select from DB record
func (messenger *Messenger) FindById(id string, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id, name FROM $1 where id = $2", messengersTable, id) //read the messenger from the database

	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(messenger.Id, messenger.Name); err != nil {
			return err
		} else {
			log.Println("Failed to extract messenger with id: ", id)
		}
		return nil
	}
	return fmt.Errorf("messenger with id=%q not found", id)
}

// Saving messenger to DB
func (messenger *Messenger) Save(tx *sql.Tx) error {
	//form the query
	query := fmt.Sprintf("INSERT INTO %s name VALUES $1 RETURNING id", messengersTable)
	//Insert into DB
	var id string
	if err := tx.QueryRow(query, messenger.Name).Scan(&id); err != nil {
		return fmt.Errorf("failed to save employer %w", err)
	}
	messenger.Id = id
	return nil
}

// Deleting the messenger from DB
func (messenger *Messenger) Delete(tx *sql.Tx) error {
	//impossible to delete messenger
	return fmt.Errorf("can't delete messenger")
}

func (messenger *Messenger) StatusUpdate(status string, tx *sql.Tx) error {
	//There is no field called status.
	return fmt.Errorf("can't update status of the messenger")
}

func (messenger *Messenger) Update(tx *sql.Tx) error {
	//impossible to update messenger
	return fmt.Errorf("can't update messenger")
}

// Getter for Id
func (messenger *Messenger) GetId() string {
	return messenger.Id
}

func GetMessengers(tx *sql.Tx) (messengerList []*Messenger, err error) {
	//Get all messengers
	query := fmt.Sprintf("SELECT id, name FROM %s", messengersTable)
	rows, err := tx.Query(query) //read the messenger from the database

	if err != nil {
		log.Println("Failed to extract messengers")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tempMesseger := NewMessenger("", "")
		if err := rows.Scan(&tempMesseger.Id, &tempMesseger.Name); err != nil {
			log.Println("Failed to parse a messenger")
			continue
		}
		messengerList = append(messengerList, tempMesseger)
	}
	log.Println(string(len(messengerList)))
	return messengerList, nil
}
