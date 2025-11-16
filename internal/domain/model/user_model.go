package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
	ProfileURL string             `bson:"profile_url,omitempty" json:"profile_url,omitempty"`
	UserRole   string             `bson:"user_role" json:"user_role"`
	CreatedDate time.Time         `bson:"created_date,omitempty" json:"created_date,omitempty"`
	CreatedBy  string             `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedDate time.Time         `bson:"updated_date,omitempty" json:"updated_date,omitempty"`
	UpdatedBy  string             `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
}

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}