package gopitman

import "github.com/gofrs/uuid"

type Essay struct {
	Id           int       `json:"-" db:"id"`
	Code_object  uuid.UUID `json:"code_object" db:"code_object"`
	Title        string    `json:"title" db:"title"`
	Address      string    `json:"address" db:"address"`
	Coordinates  string    `json:"coordinates" db:"coordinates"`
	Descriptions string    `json:"descriptions" db:"descriptions"`
}
