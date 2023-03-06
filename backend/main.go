package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// CONSTANTS

const appStatus string = "ONLINE"
const port string = ":9001"

// MAIN

func main() {
	fmt.Println("Server started and listening on port " + port)

	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", putRequest).Methods("PUT")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(port, nil))

	// open database connection
	db, err := sql.Open("mysql", "root:changeme@tcp(127.0.0.1:3306)/activitytracker_db")
	// handle the error if one occurs when opening the connection
	checkError(err)
	// defer the close till after the main function has finished executing
	defer db.Close()

	createQuery, err := db.Query("INSERT INTO activities VALUES (1, 'Karol Pawlak', 'Running', 9000, 21.1)")
	checkError(err)

	defer createQuery.Close()
}

// REQUESTS

func getRequest(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println("\nGET request received on ", t.Format(time.RFC3339), w)
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPOST request received", w)
}

func putRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPUT request received", w)
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nDELETE request received", w)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		//panic(err.Error())
		//log.Fatal(err)
	}
}

// v1 -----------------------------------------------------------------------

// func main() {
// 	a := AppConfig{}
// 	a.port = ":9001"
// 	a.Init()
// 	a.Run()
// }

// REQUESTS

// func basicRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, appStatus)
// 	a := Activity{"Karol", "Running", 2400, 8.0}
// 	fmt.Println(a.calculatePace())
// }

// DATABASE STUFF

// type AppConfig struct {
// 	DB *sql.DB
// 	Port string
// 	Router *mux.Router
// }

// func (a *AppConfig) Run() {
// 	fmt.Println("Server started and listening on port ", a.Port)
// 	log.Fatal(http.ListenAndServe(a.Port, a.Router))
// }

// func (a *AppConfig) InitializeRoutes() {
// 	fmt.Println("Routes initialized")
// 	a.Router.HandleFunc("/", basicRequest)
// }

// func (a *AppConfig) Init() {
// 	a.DB, err := sql.Open("sqlite3", "../../test.db")
// 	checkError(err)

//  a.router = mux.NewRouter()
//  a.InitializeRoutes()
// }

// func getRows() {
// 	rows, err := db.Query("SELECT * FROM activities")
// 	checkError(err)
// 	defer rows.Close()

// 	for rows.Next() {
// 		var a Activity
// 		rows.Scan(&a.userName, &a.activityType, &a.activityLength, a.distance)
// 		fmt.Println("Activity: ", a.activityType, " By: ", a.userName, " Distance: ", a.distance)
// 	}
// }

// ERROR CHECKING
