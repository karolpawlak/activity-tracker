package main

import (
	"database/sql"
	"fmt"
)

type Activity struct {
	ID             int     `json:"activity_id"`
	userName       string  `json:"user_name"`
	activityType   string  `json:"activity_type"`
	activityLength int     `json:"activity_length"` // in seconds
	distance       float32 `json:"distance"`        // in kilometers

}

func (a *Activity) calculatePace() float32 {
	return (float32(a.activityLength) / a.distance) / 60 // return result in minutes
}

// DB QUERIES

func (a *Activity) createNewActivity(db *sql.DB) error {
	response, err := db.Exec("INSERT INTO activities(user_name, activity_type, activity_length, distance) VALUES (?, ?, ?, ?)", a.userName, a.activityType, a.activityLength, a.distance)
	checkError(err)

	id, err := response.LastInsertId()
	checkError(err)
	a.ID = int(id)

	return nil
}

func (a *Activity) getSingleActivity(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM activities WHERE id = ?", a.ID).Scan(&a.ID, &a.userName, &a.activityType, &a.activityLength, &a.distance)
}

func getAggregateActivities(db *sql.DB) ([]Activity, error) {
	rows, err := db.Query("SELECT * FROM activities")
	checkError(err)
	defer rows.Close()

	activities := []Activity{}

	// for each row, scan the result into an activity composite object
	for rows.Next() {
		var activity Activity
		err = rows.Scan(&activity.ID, &activity.userName, &activity.activityType, &activity.activityLength, &activity.distance)
		checkError(err)

		fmt.Println("ID: ", activity.ID, "Activity: ", activity.activityType, " By: ", activity.userName, " Distance: ", activity.distance)
		activities = append(activities, activity)
	}

	return activities, err
}
