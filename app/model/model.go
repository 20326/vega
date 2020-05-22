package model

import "time"

type (
	// Model represents meta data of entity.
	Model struct {
		ID        uint64     `gorm:"primary_key" json:"id"`
		CreatedAt time.Time  `sql:"index" json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
