package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	// open database connection
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
	app.Router.HandleFunc("/single/{id}", app.getRequestSingle).Methods("GET")
	app.Router.HandleFunc("/all", app.getRequestAll).Methods("GET")
	app.Router.HandleFunc("/", postRequest).Methods("POST")
	app.Router.HandleFunc("/", putRequest).Methods("PUT")
	app.Router.HandleFunc("/", deleteRequest).Methods("DELETE")
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

	var activity Activity
	activity.ID = int(id)

	// execute query and respond to the client
	err = activity.getSingleActivity(app.Database)
	if err != nil {
		fmt.Printf("getRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	fmt.Println("ID: ", activity.ID, "Activity: ", activity.activityType, " By: ", activity.userName, " Distance: ", activity.distance)
	respondWithJSON(w, http.StatusOK, activity)

}

func (app *AppConfig) getRequestAll(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println("\nGET request received on ", t.Format(time.RFC3339), w)

	// execute query and respond to the client
	activities, err := getAllActivities(app.Database)
	if err != nil {
		fmt.Printf("getRequest error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, activities)
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPOST request received", w)

	activity := Activity{1, "Karol", "Running", 2400, 8.0}
	fmt.Println(activity.calculatePace())
}

func putRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPUT request received", w)
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nDELETE request received", w)
}

// func basicRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, appStatus)
// 	a := Activity{"Karol", "Running", 2400, 8.0}
// 	fmt.Println(a.calculatePace())
// }

// createQuery, err := app.Database.Query("INSERT INTO activities VALUES (2, 'Karol Pawlak', 'Running', 9000, 21.1)")
// checkError(err)
// defer createQuery.Close()

// ERROR CHECKING

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		//panic(err.Error()) // proper error handling instead of panic in your app
		//log.Fatal(err)
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
