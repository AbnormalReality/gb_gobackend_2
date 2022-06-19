package corporate

import "github.com/google/uuid"


type Corporate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	
}