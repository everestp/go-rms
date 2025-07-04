package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string            `json:"name" validate:"required,min=2,max=100"`
	Category   string            `json:"category" validate:"required,min=2,max=100"`
	Start_Date *time.Time         `json:"start_date" validate:"required"`
	End_Date   *time.Time         `json:"end_date" validate:"required"`
	Created_at 	time.Time         `json:"created_at"`
	Updated_at 	time.Time         `json:"updated_at"`
	Menu_id    string            `json:"menu_id" validate:"required"`
}
