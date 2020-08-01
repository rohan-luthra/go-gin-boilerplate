package models

import "github.com/go-playground/validator/v10"

type User struct {
	ID  primitive.ObjectID`bson:"_id,omitempty"`
	Username string `bson:"username,omitempty" validate:"required"`
}

func (u *User) validate() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return err
	}
}