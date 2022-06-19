package communities

import "github.com/google/uuid"


type Communities struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	
}