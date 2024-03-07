package models

type Posts struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
}
