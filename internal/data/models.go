package data

// Scrum represents scrum data
type Scrum struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	ProjectID  int     `json:"project_id"`
	WorkDone   string  `json:"work_done"`
	Todos      string  `json:"todos"`
	HoursSpent float32 `json:"hours_spent"`
	Date       string  `json:"date"`
	CreatedBy  int     `json:"created_by"`
	UpdatedBy  int     `json:"updated_by"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
