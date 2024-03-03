package models

type SchedulePost struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Username    string `json:"username"`
	ScheduledAt string `json:"scheduledAt"`
}
