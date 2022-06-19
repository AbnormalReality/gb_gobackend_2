package projects

import "github.com/google/uuid"


type Projects struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
}