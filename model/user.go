package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User holds metadata about a user.
type (
	User struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		CreateAt  time.Time     `json:"-" bson:"createAt"`
		Name      string        `json:"name" bson:"name,omitempty" validate:"nonzero"`
		Email     string        `json:"email" bson:"email,omitempty" validate:"nonzero"`
		BirthDate string        `json:"birthDate" bson:"birthDate,omitempty" validate:"len=10"`
		Gender    string        `json:"gender" bson:"gender,omitempty" validate:"regexp=(^male$|^female$)"`
		Hospital  string        `json:"hospital" bson:"hospital,omitempty" validate:"nonzero"`
		Weight    int           `json:"weight,omitempty" bson:"weight,omitempty"`
		Height    int           `json:"height,omitempty" bson:"height,omitempty"`
	}
)
