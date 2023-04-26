package main

import (
	"database/sql"
	"fmt"
)

type Activity struct {
	ID             int     `json:"activity_id"`
	UserName       string  `json:"user_name"`
	ActivityType   string  `json:"activity_type"`
	ActivityLength int     `json:"activity_length"` // in seconds
	Distance       float32 `json:"distance"`        // in kilometers

}

func (a *Activity) calculatePace() float32 {
	return (float32(a.ActivityLength) / a.Distance) / 60 // return result in minutes
}

// DB QUERIES

func (a *Activity) createNewActivity(db *sql.DB) error {
	//response, err := db.Exec("INSERT INTO activities(user_name, activity_type, activity_length, distance) VALUES (?, ?, ?, ?)", a.userName, a.activityType, a.activityLength, a.distance)
	createQuery, err := db.Prepare("INSERT INTO activities(user_name, activity_type, activity_length, distance) VALUES (?, ?, ?, ?)")
	checkError(err)

	response, err := createQuery.Exec(a.UserName, a.ActivityType, a.ActivityLength, a.Distance)
	checkError(err)

	id, err := response.LastInsertId()
	checkError(err)
	a.ID = int(id)

	return nil
}

func (a *Activity) getSingleActivity(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM activities WHERE id = ?", a.ID).Scan(&a.ID, &a.UserName, &a.ActivityType, &a.ActivityLength, &a.Distance)
}

func getAggregateActivities(db *sql.DB) ([]Activity, error) {
	rows, err := db.Query("SELECT * FROM activities")
	checkError(err)
	defer rows.Close()

	activities := []Activity{}

	// for each row, scan the result into an activity composite object
	for rows.Next() {
		var activity Activity
		err = rows.Scan(&activity.ID, &activity.UserName, &activity.ActivityType, &activity.ActivityLength, &activity.Distance)
		checkError(err)

		fmt.Println("ID: ", activity.ID, "Activity: ", activity.ActivityType, " By: ", activity.UserName, " Distance: ", activity.Distance)
		activities = append(activities, activity)
	}

	return activities, err
}
