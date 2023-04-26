package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// CONSTANTS

const appStatus string = "ONLINE"
const port string = ":9001"

type AppConfig struct {
	Database *sql.DB
	Port     string
	Router   *mux.Router
}

func (app *AppConfig) Run() {
	fmt.Println("Server started and listening on port ", app.Port)
	log.Fatal(http.ListenAndServe(app.Port, app.Router))
}

func (app *AppConfig) Init() {
	// open database connection - N.B returns a pointer to the database
	db, err := sql.Open("mysql", "root:changeme@tcp(127.0.0.1:3306)/activitytracker_db")
	app.Database = db

	// handle the error if one occurs when opening the connection
	checkError(err)

	// create a router and initialize routes
	app.Router = mux.NewRouter()
	app.InitializeRoutes()
}

func (app *AppConfig) InitializeRoutes() {
	fmt.Println("Routes initialized")
	app.Router.HandleFunc("/activities/{id}", app.getRequestSingle).Methods("GET")
	app.Router.HandleFunc("/activities", app.getRequestAggregate).Methods("GET")
	app.Router.HandleFunc("/activities/new", app.postRequest).Methods("POST")
	app.Router.HandleFunc("/activities/update/{id}", app.putRequest).Methods("PUT")
	app.Router.HandleFunc("/activities/delete/{id}", app.deleteRequest).Methods("DELETE")
	http.Handle("/", app.Router)
}

// MAIN

func main() {
	app := AppConfig{}
	app.Port = port
	app.Init()
	app.Run()

	// defer the close till after the main function has finished executing
	defer app.Database.Close()
}

// REQUESTS

func (app *AppConfig) getRequestSingle(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println("\nGET request received on ", t.Format(time.RFC3339), w)

	// get variable from the request and populate in activity object
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	checkError(err)

	var a Activity
	a.ID = int(id)

	// execute query and respond to the client
	err = a.getSingleActivity(app.Database)
	if err != nil {
		fmt.Printf("getRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, a)
}

func (app *AppConfig) getRequestAggregate(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println("\nGET request received on ", t.Format(time.RFC3339), w)

	// execute query and respond to the client
	activities, err := getAggregateActivities(app.Database)
	if err != nil {
		fmt.Printf("getRequestAggregate error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, activities)
}

func (app *AppConfig) postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPOST request received", w)

	requestBody, _ := ioutil.ReadAll(r.Body)

	var a Activity
	json.Unmarshal(requestBody, &a)

	err := a.createNewActivity(app.Database)
	if err != nil {
		fmt.Printf("postRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusCreated, a)
}

func (app *AppConfig) putRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPUT request received", w)

	// get variable from the request and populate in activity object
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	checkError(err)

	var a Activity
	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &a)
	a.ID = int(id)

	// execute query and respond to the client
	err = a.updateActivity(app.Database)
	if err != nil {
		fmt.Printf("putRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, a)
}

func (app *AppConfig) deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nDELETE request received", w)

	// get variable from the request and populate in activity object
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	checkError(err)

	var a Activity
	a.ID = int(id)

	// execute query and respond to the client
	err = a.deleteActivity(app.Database)
	if err != nil {
		fmt.Printf("deleteRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusNoContent, a)
}

// ERROR CHECKING

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		//panic(err.Error()) // proper error handling instead of panic in your app
		//log.Fatal(err)
		//return err
	}
}

// HELPER FUNCTIONS

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message}) // package the error and call respondWithJSON function
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// create a response
	response, _ := json.Marshal(payload)

	// set headers and write a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
