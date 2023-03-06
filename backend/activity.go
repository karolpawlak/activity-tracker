package main

type Activity struct {
	ID             int     `json:"id"`
	userName       string  `json:"user_name"`
	activityType   string  `json:"activity_type"`
	activityLength int     `json:"activity_length"` // in seconds
	distance       float32 `json:"distance"`        // in kilometers

}

func (a Activity) calculatePace() float32 {
	return (float32(a.activityLength) / a.distance) / 60 // return result in minutes
}
