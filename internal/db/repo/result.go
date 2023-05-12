package repo

import "time"

type Account struct {
	AccountNumber string    `json:"account_number"`
	UserID        int       `json:"user_id"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

type Transfer struct {
	ID          int
	FromAccount string
	ToAccount   string
	Amount      float64
	CreatedAt   time.Time
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
