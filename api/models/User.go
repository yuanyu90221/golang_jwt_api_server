package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"

	"golang.org/x/crypto/bcrypt"
)

type validator interface {
	validate(u *User) error
}

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

type updateAction struct {
}

func (a *updateAction) validate(u *User) error {
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

type loginAction struct {
}

func (a *loginAction) validate(u *User) error {
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

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Passwd    string    `gorm:"size:100;not null;" json:"passwd"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

func VerifyPasswd(hashedPasswd, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passwd))
}

func (u *User) BeforeSave() error {
	hashedPasswd, err := Hash(u.Passwd)
	if err != nil {
		return err
	}
	u.Passwd = string(hashedPasswd)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(v validator) error {
	return v.validate(u)
}
