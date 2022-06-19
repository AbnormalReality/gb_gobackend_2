package organiztions

import "github.com/google/uuid"


type Organizations struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`

}