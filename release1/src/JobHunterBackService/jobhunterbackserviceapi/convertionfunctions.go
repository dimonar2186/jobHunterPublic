package jobhunterbackserviceapi

import (
	"log"
	"time"

	"JobHunterBackService/dbentity/country"
	"JobHunterBackService/dbentity/currency"
	"JobHunterBackService/dbentity/employer"
	"JobHunterBackService/dbentity/messenger"

	"github.com/google/uuid"
)

// convertCountriesconverts a slice of country data to the response format
func convertCountries(countries []*country.Country) []Country {
	responseCountries := make([]Country, 0, len(countries))
	for _, country := range countries {
		parsedUUID, err := uuid.Parse(country.Id)
		if err != nil {
			log.Println("Error in UUID conversion:", err)
			continue // Skip this country if there's an error
		}
		tempCountry := Country{
			Code: &country.Code,
			Id:   &parsedUUID,
			Name: &country.Name,
		}
		responseCountries = append(responseCountries, tempCountry)
	}
	return responseCountries
}

// convertCountryconvers a single country data to the response format
func convertCountry(dbcountry country.Country) Country {
	parsedUUID, err := uuid.Parse(dbcountry.Id)
	if err != nil {
		log.Println("Error in UUID conversion:", err)
	}
	tempCountry := Country{
		Code: &dbcountry.Code,
		Id:   &parsedUUID,
		Name: &dbcountry.Name,
	}
	return tempCountry
}

// convertCurrenciesconverts a slice of currency data to the response format
func convertCurrencies(currencies []*currency.Currency) []Currency {
	responseCurrencies := make([]Currency, 0, len(currencies))
	for _, currency := range currencies {
		parsedUUID, err := uuid.Parse(currency.Id)
		if err != nil {
			log.Println("Error in UUID conversion:", err)
			continue // Skip this country if there's an error
		}
		tempCurrency := Currency{
			Id:        &parsedUUID,
			IsoCode:   &currency.IsoCode,
			Name:      &currency.Name,
			ShortName: &currency.ShortName,
		}
		responseCurrencies = append(responseCurrencies, tempCurrency)
	}
	return responseCurrencies
}

// convertCurrencyconvers a single currency data to the response format
func convertCurrency(dbcurrency currency.Currency) Currency {
	parsedUUID, err := uuid.Parse(dbcurrency.Id)
	if err != nil {
		log.Println("Error in UUID conversion:", err)
	}
	tempCurrency := Currency{
		Id:        &parsedUUID,
		IsoCode:   &dbcurrency.IsoCode,
		Name:      &dbcurrency.Name,
		ShortName: &dbcurrency.ShortName,
	}
	return tempCurrency
}

// convertMessengersconverts a slice of messengers data to the response format
func convertMessengers(messengers []*messenger.Messenger) []Messenger {
	responseMessengers := make([]Messenger, 0, len(messengers))
	for _, messenger := range messengers {
		parsedUUID, err := uuid.Parse(messenger.Id)
		if err != nil {
			log.Println("Error in UUID conversion:", err)
			continue // Skip this messenger if there's an error
		}
		tempMessenger := Messenger{
			Id:   &parsedUUID,
			Name: &messenger.Name,
		}
		responseMessengers = append(responseMessengers, tempMessenger)
	}
	return responseMessengers
}

func convertEmployerBaseServerToDB(employerBase EmployerBase) *employer.Employer {
	employerDB := employer.NewEmployer("", "", time.Time{})
	employerDB.Name = employerBase.Name
	return employerDB
}

func convertLocationsServerToDBCountries(locations *[]Country) (dbcountries []*country.Country) {
	for _, location := range *locations {
		tempCountry := country.Country{
			Code: *location.Code,
			Id:   (*location.Id).String(),
			Name: *location.Name,
		}
		dbcountries = append(dbcountries, &tempCountry)
	}
	return dbcountries
}

func convertEmployerResponseDBToServer(employerDb *employer.Employer, countriesDb []*country.Country) *EmployerResponse {
	parsedUUID, err := uuid.Parse(employerDb.Id)
	if err != nil {
		log.Println("Error in UUID conversion:", err)
	}
	employerResponse := EmployerResponse{
		Id:           parsedUUID,
		CreationDate: &employerDb.CreationDate,
		Name:         employerDb.Name,
		UpdateDate:   &time.Time{},
	}

	var locations []Country
	for _, countryDb := range countriesDb {
		location := convertCountry(*countryDb)
		locations = append(locations, location)
	}
	employerResponse.Locations = &locations
	return &employerResponse
}
