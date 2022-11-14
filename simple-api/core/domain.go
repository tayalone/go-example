package core

import "time"

/*Todo is Definition of Domain in Application*/
type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createAt"`
}
