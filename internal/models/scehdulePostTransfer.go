package models

type ScehdulePostTransfer struct {
	ID       string `json:"id" bson:"_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
}
