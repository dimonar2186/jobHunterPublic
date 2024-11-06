// implementation of the server

package jobhunterbackserviceapi

import (
	"log"
	"net/http"

	//"sync"

	//"github.com/go-chi/chi/v5"
	//"github.com/gorilla/mux"
	"encoding/json"

	"JobHunterBackService/dbcrud"
	"JobHunterBackService/dbentity/country"
	"JobHunterBackService/dbentity/currency"
	"JobHunterBackService/dbentity/employer"
	"JobHunterBackService/dbentity/messenger"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// sendPetStoreError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
/*func sendPetStoreError(w http.ResponseWriter, code int, message string) {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(petErr)
}*/

func (Server) GetCountries(w http.ResponseWriter, r *http.Request) {

	log.Println("countries requested")
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
	}
	defer tx.Rollback() //ready for rollback

	//Get countries from DB
	countries, err := country.GetCountries(tx)
	if err != nil {
		log.Println("failed to get countries %w", err)
		sendJobHunterBackServiceError(w, http.StatusNotFound, "Countries are not found")
		return
	}

	// Convert to response format
	response := struct {
		Countries []Country `json:"countries"`
	}{
		Countries: convertCountries(countries),
	}
	response.Countries = convertCountries(countries)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
}

// Get specific country
// (GET /countries/{countryId})
func (Server) GetCountriesCountryId(w http.ResponseWriter, r *http.Request, countryId CountryId) {

	log.Printf("country with id=%s requested", countryId)
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
	}
	defer tx.Rollback() //ready for rollback

	//Get country from DB by it's id
	var tempCountry country.Country
	log.Println("temp country created")
	if err := tempCountry.FindById(countryId.String(), tx); err != nil {
		log.Println("failed to get country with id = ", countryId.String(), " %w", err)
		sendJobHunterBackServiceError(w, http.StatusNotFound, "Country is not found")
		return
	}

	//form and send the response in json format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(convertCountry(tempCountry)); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

}

// Get the list of currencies
// (GET /currencies)
func (Server) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	log.Println("currencies requested")
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
	}
	defer tx.Rollback() //ready for rollback

	//Get countries from DB
	currencies, err := currency.GetCurrencies(tx)
	if err != nil {
		log.Println("failed to get currencies %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	// Convert to response format
	response := struct {
		Currencies []Currency `json:"currencies"`
	}{
		Currencies: convertCurrencies(currencies),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
}

// Get the list of currencies
// (GET /currencies/{currencyId})
func (Server) GetCurrenciesCurrencyId(w http.ResponseWriter, r *http.Request, currencyId openapi_types.UUID) {

	log.Printf("currency with id=%s requested", currencyId)
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
	}
	defer tx.Rollback() //ready for rollback

	//Get currency from DB by it's id
	var tempCurrency currency.Currency
	log.Println("temp currency created")
	if err := tempCurrency.FindById(currencyId.String(), tx); err != nil {
		log.Println("failed to get currency with id = ", currencyId.String(), " %w", err)
		sendJobHunterBackServiceError(w, http.StatusNotFound, "Currency is not found")
		return
	}

	//form and send the response in json format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(convertCurrency(tempCurrency)); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

}

// Get the list of messangers
// (GET /messangers)
func (Server) GetMessangers(w http.ResponseWriter, r *http.Request) {
	log.Println("messengers requested")
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
	}
	defer tx.Rollback() //ready for rollback

	//Get messengers from DB
	messengers, err := messenger.GetMessengers(tx)
	if err != nil {
		log.Println("failed to get messengers %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	// Convert to response format
	response := struct {
		ExportMessengers []Messenger `json:"messangers"`
	}{
		ExportMessengers: convertMessengers(messengers),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
}

// Delete specific HRManager
// (DELETE /user/HRMAnagers/{HRManagerId})
func (Server) DeleteUserHRMAnagersHRManagerId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId) {
}

// Get specific HRManager of the vacancy
// (GET /user/HRMAnagers/{HRManagerId})
func (Server) GetUserHRMAnagersHRManagerId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId) {
}

// Update specific HRManager of the vacancy
// (PATCH /user/HRMAnagers/{HRManagerId})
func (Server) PatchUserHRMAnagersHRManagerId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId) {
}

// Creates a new HRManager
// (POST /user/HRManagers)
func (Server) PostUserHRManagers(w http.ResponseWriter, r *http.Request) {
}

// Returns list of contacts of the HRManager
// (GET /user/HRManagers/{HRManagerId}/contacts)
func (Server) GetUserHRManagersHRManagerIdContacts(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId) {
}

// creates a new HRManager contact
// (POST /user/HRManagers/{HRManagerId}/contacts)
func (Server) PostUserHRManagersHRManagerIdContacts(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId) {
}

// deletes the existing HRManager's contact
// (DELETE /user/HRManagers/{HRManagerId}/contacts/{contactId})
func (Server) DeleteUserHRManagersHRManagerIdContactsContactId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId, contactId ContactId) {
}

// Returns specific contact of the HRManager
// (GET /user/HRManagers/{HRManagerId}/contacts/{contactId})
func (Server) GetUserHRManagersHRManagerIdContactsContactId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId, contactId ContactId) {
}

// updates a new HRManager contact
// (PATCH /user/HRManagers/{HRManagerId}/contacts/{contactId})
func (Server) PatchUserHRManagersHRManagerIdContactsContactId(w http.ResponseWriter, r *http.Request, hrManagerId HRManagerId, contactId ContactId) {
}

// Create a new applicationStage
// (POST /user/applicationStages)
func (Server) PostUserApplicationStages(w http.ResponseWriter, r *http.Request) {
}

// update applicationStage
// (DELETE /user/applicationStages/{applicationStageId})
func (Server) DeleteUserApplicationStagesApplicationStageId(w http.ResponseWriter, r *http.Request, applicationStageId ApplicationStageId) {
}

// update status of applicationStage
// (PATCH /user/applicationStages/{applicationStageId})
func (Server) PatchUserApplicationStagesApplicationStageId(w http.ResponseWriter, r *http.Request, applicationStageId ApplicationStageId) {
}

// Create a new employer
// (POST /user/employers)
func (Server) PostUserEmployers(w http.ResponseWriter, r *http.Request) {

	// Get a EmployerBase from requestBody
	var newEmployer EmployerBase
	if err := json.NewDecoder(r.Body).Decode(&newEmployer); err != nil {
		sendJobHunterBackServiceError(w, http.StatusBadRequest, "Invalid format for newEmployer")
		return
	}

	log.Println("new employer provided")
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
	defer tx.Rollback() //ready for rollback

	//try to save employer
	employerDB := convertEmployerBaseServerToDB(newEmployer)
	if err = employerDB.Save(tx); err != nil {
		log.Printf("failed to save emp %s", newEmployer.Name)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	//try to save locations
	countriesDb := convertLocationsServerToDBCountries(newEmployer.Locations)
	for _, countryDb := range countriesDb {
		if err = countryDb.SaveLocation(employerDB, tx); err != nil {
			log.Printf("failed to save emp location")
			sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
			return
		}
	}

	// if everything is fine, then commit the transaction
	tx.Commit()

	// create and send a response
	response := convertEmployerResponseDBToServer(employerDB, countriesDb)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
}

// Deletes the existing employer
// (DELETE /user/employers/{employerId})
func (Server) DeleteUserEmployersEmployerId(w http.ResponseWriter, r *http.Request, employerId EmployerId) {
	log.Println("need to delete employer")

	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
	defer tx.Rollback() //ready for rollback

	//Choose the existing employer
	//Let's find an employerDB with employerId. If not found then send 404
	var employerDB employer.Employer
	if err := employerDB.FindById(employerId.String(), tx); err != nil {
		log.Printf("failed to update emp, because this id is not found %s %+v", employerId.String(), err)
		sendJobHunterBackServiceError(w, http.StatusNotFound, "not found")
		return
	}

	//delete all locations for employer
	if err = country.DeleteAllCountries(&employerDB, tx); err != nil {
		log.Printf("failed to delete locations for emp %s", employerDB.Id)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	//try to delete employer
	if err = employerDB.Delete(tx); err != nil {
		log.Printf("failed to delete employer with %s : %+v", employerDB.Id, err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	// if everything is fine, then commit the transaction
	tx.Commit()

	w.WriteHeader(http.StatusOK)

}

// Updates the existing employer
// (PATCH /user/employers/{employerId})
func (Server) PatchUserEmployersEmployerId(w http.ResponseWriter, r *http.Request, employerId EmployerId) {
	// Get a EmployerBase from requestBody
	var patchedEmployer EmployerBase
	if err := json.NewDecoder(r.Body).Decode(&patchedEmployer); err != nil {
		sendJobHunterBackServiceError(w, http.StatusBadRequest, "Invalid format for patchedEmployer")
		return
	}

	log.Println("employer's update is provided")
	//Open transaction
	tx, err := dbcrud.OpenTransaction()
	if err != nil {
		log.Println("failed to open transaction %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
	defer tx.Rollback() //ready for rollback

	//Let's find an employerDB with employerId. If not found then send 404
	var employerDB employer.Employer
	if err := employerDB.FindById(employerId.String(), tx); err != nil {
		log.Printf("failed to update emp, because this id is not found %s", patchedEmployer.Name)
		sendJobHunterBackServiceError(w, http.StatusNotFound, "not found")
		return
	}

	//try to update employer
	employerDB = *convertEmployerBaseServerToDB(patchedEmployer)
	employerDB.Id = employerId.String()
	if err = employerDB.Update(tx); err != nil {
		log.Printf("failed to update emp %s", patchedEmployer.Name)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	//here i didn't want to think. so i decided to remove all locations for employer and create new locations.
	if err = country.DeleteAllCountries(&employerDB, tx); err != nil {
		log.Printf("failed to delete locations for emp %s", employerDB.Id)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}

	//try to save locations
	countriesDb := convertLocationsServerToDBCountries(patchedEmployer.Locations)
	for _, countryDb := range countriesDb {
		if err = countryDb.SaveLocation(&employerDB, tx); err != nil {
			log.Printf("failed to save emp location")
			sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
			return
		}
	}

	// if everything is fine, then commit the transaction
	tx.Commit()

	// create and send a response
	response := convertEmployerResponseDBToServer(&employerDB, countriesDb)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode response %w", err)
		sendJobHunterBackServiceError(w, http.StatusGatewayTimeout, "something went wrong")
		return
	}
}

// Returns a list of jobSearchingProcesses associated with current user
// (GET /user/jobSearchingProcesses)
func (Server) GetUserJobSearchingProcesses(w http.ResponseWriter, r *http.Request) {
}

// create a new jobSearchingProcess
// (POST /user/jobSearchingProcesses)
func (Server) PostUserJobSearchingProcesses(w http.ResponseWriter, r *http.Request) {
}

// delete the existing jobSearchingProcess
// (DELETE /user/jobSearchingProcesses/{jobSearchingProcessId})
func (Server) DeleteUserJobSearchingProcessesJobSearchingProcessId(w http.ResponseWriter, r *http.Request, jobSearchingProcessId openapi_types.UUID) {
}

// Get properties of a specified jobSearchingProcess
// (GET /user/jobSearchingProcesses/{jobSearchingProcessId})
func (Server) GetUserJobSearchingProcessesJobSearchingProcessId(w http.ResponseWriter, r *http.Request, jobSearchingProcessId string) {
}

// update existing jobSearchingProcess
// (PATCH /user/jobSearchingProcesses/{jobSearchingProcessId})
func (Server) PatchUserJobSearchingProcessesJobSearchingProcessId(w http.ResponseWriter, r *http.Request, jobSearchingProcessId string) {
}

// create a new vacancy
// (POST /user/vacancies)
func (Server) PostUserVacancies(w http.ResponseWriter, r *http.Request) {
}

// delete specific vacancy
// (DELETE /user/vacancies/{vacancyIdParam})
func (Server) DeleteUserVacanciesVacancyIdParam(w http.ResponseWriter, r *http.Request, vacancyIdParam VacancyIdParam) {
}

// Get specified vacancy
// (GET /user/vacancies/{vacancyIdParam})
func (Server) GetUserVacanciesVacancyIdParam(w http.ResponseWriter, r *http.Request, vacancyIdParam VacancyIdParam) {
}

// Update the existing vacancy
// (PATCH /user/vacancies/{vacancyIdParam})
func (Server) PatchUserVacanciesVacancyIdParam(w http.ResponseWriter, r *http.Request, vacancyIdParam VacancyIdParam) {
}

type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

func sendJobHunterBackServiceError(w http.ResponseWriter, code int, message string) {
	jobHunterErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jobHunterErr)
}
