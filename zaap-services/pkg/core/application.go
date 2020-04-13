package core

import "time"

type ApplicationState string

const (
	ApplicationStateUnknown  ApplicationState = "unknown"
	ApplicationStateStopped                   = "stopped"
	ApplicationStateStarting                  = "starting"
	ApplicationStateRunning                   = "running"
)

type Application struct {
	ID          string            `json:"id"`
	Environment map[string]string `json:"environment"`
	Image       string            `json:"image"`
	Name        string            `json:"name"`
	State       ApplicationState  `json:"state"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	UserID      string            `json:"user_id"`
}

type ApplicationStateChanged struct {
	ApplicationID string           `json:"application_id"`
	State         ApplicationState `json:"state"`
}

type DeploymentPayload struct {
	SchedulerToken string      `json:"scheduler_token"`
	Application    Application `json:"application"`
}
