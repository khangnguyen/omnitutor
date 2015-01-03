package models

import (
  "time"
)

type UserAccount struct {
  Id              int64
  Nickname        string
  Email           string
  Username        string
  Password        string
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