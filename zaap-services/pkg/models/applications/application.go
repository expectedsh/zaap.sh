package application

import "time"

type Application struct {
	ID          string            `json:"id"`
	Environment map[string]string `json:"environment"`
	Image       string            `json:"image"`
	Name        string            `json:"name"`
	State       string            `json:"state"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	UserID      string            `json:"user_id"`
}

type DeploymentPayload struct {
	SchedulerToken string      `json:"scheduler_token"`
	Application    Application `json:"application"`
}
