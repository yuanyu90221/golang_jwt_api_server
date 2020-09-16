package models

import (
	"errors"

	"github.com/badoux/checkmail"
)

//Validator interface
type Validator interface {
	validate(u *User) error
}

//DefaultAction struct
type DefaultAction struct {
}

func (a *DefaultAction) validate(u *User) error {
	if u.Nickname == "" {
		return errors.New("Required NickName")
	}
	if u.Passwd == "" {
		return errors.New("Required Password")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

//UpdateAction struct
type UpdateAction struct {
}

func (a *UpdateAction) validate(u *User) error {
	if u.Nickname == "" {
		return errors.New("Required NickName")
	}
	if u.Passwd == "" {
		return errors.New("Required Password")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

//LoginAction struct
type LoginAction struct {
}

func (a *LoginAction) validate(u *User) error {
	if u.Passwd == "" {
		return errors.New("Required Password")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}
