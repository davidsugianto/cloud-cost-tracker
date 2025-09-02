package budget

import "time"

type Budget struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
