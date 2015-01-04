package models

import (
  "time"
  "github.com/revel/revel"
)

type UserAccount struct {
  Id              int64
  Nickname        string
  Email           string
  Username        string
  Password        string
  PasswordConfirm string `sql:"-"`
  Gender          string
  AvatarFilename  string
  InviterId       int64
  CreatedAt       time.Time `sql:"DEFAULT:current_timestamp"`
  UpdatedAt       time.Time `sql:"DEFAULT:current_timestamp"`
  LastLoginAt     time.Time `sql:"DEFAULT:current_timestamp"`
  IsArchived      bool `sql:"DEFAULT:false"`
  Locale          string `sql:"DEFAULT:'en'"`
}

func (u UserAccount) TableName() string {
  return "user_account"
}

func (user *UserAccount) Validate(v *revel.Validation) {
  v.Required(user.Email).
          Message("The field is required")
  v.Email(user.Email).
       Message("Email format is invalid")
  v.Required(user.Username).
          Message("The field is required")
  v.Required(user.Password).
          Message("The field is required")
  v.Required(user.PasswordConfirm == user.Password).
          Message("The passwords do not match")
}