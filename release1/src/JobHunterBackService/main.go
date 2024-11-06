package main

import (
	//"fmt"
	"log"
	"os"

	//"time"

	//"JobHunterBackService/dbcrud"
	//"JobHunterBackService/dbentity/country"
	//"JobHunterBackService/dbentity/jobsearchingprocess"
	"JobHunterBackService/jobhunterbackserviceapi"

	_ "github.com/lib/pq"
	//_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	"github.com/go-chi/chi/v5"

	//"flag"
	//"net"
	"net/http"
	//middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	//Open a file for logs
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Route logs in the file
	log.SetOutput(file)

	//tx, err := dbcrud.OpenTransaction()

	//jsp := jobsearchingprocess.NewJobSearchingProcess("1", "Software Developer", "00000000-0000-0000-0000-000000000000", "active", "00000000-0000-0000-0000-000000000000", 3000, 5000, time.Now(), time.Now(), time.Time{}, false, []string{"Developer", "SU"})
	//("1", "Software Developer", 3000, 5000, []string{"Developer"}, "user123", time.Now(), time.Now(), "Active", false)

	/*jsp := jobsearchingprocess.Jobsearchingprocess{
	Id:                   "1",
	Name:                 "Software Developer",
	MinimumMonthlySalary: 3000,
	MaximumMonthlySalary: 5000,
	Positions:            []string{"Developer"},
	User:                 "user123",
	CreationDate:         time.Now(),
	UpdateDate:           time.Now(),
	Status:               "Active",
	IsDeleted:            false}

	err = jsp.Save(tx)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("saved ", jsp.Id)*/

	/*var tempCountry []*country.Country
	tempCountry, err = country.GetAllLocations(jsp, tx)
	if tempCountry == nil {
		fmt.Println("")
	}
	fmt.Println(err)*/
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := jobhunterbackserviceapi.NewServer()

	r := chi.NewRouter()

	// get an `http.Handler` that we can use
	h := jobhunterbackserviceapi.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "127.0.0.1:999",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

}

/*func connectDB() (*sql.DB, error) {
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PWD")
	//dbName := os.Getenv("DB_NAME")

	dbUser := "jobHunterBasic_su"
	dbPassword := "A211186dssu"
	dbName := "jobHunterBasic"

	//if dbUser == "" || dbPassword == "" || dbName == "" {
	//	return nil, errors.New("database credentials are not set")
	//}
	if dbUser == "" {
		return nil, errors.New("database cred1 are not set")
	}
	if dbPassword == "" {
		return nil, errors.New("database cred2 are not set")
	}
	if dbName == "" {
		return nil, errors.New("database cred3 are not set")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}*/
