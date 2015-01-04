package controllers

import (
  "github.com/revel/revel"
  "omnitutor/app/models"
)

type Account struct {
  *revel.Controller
}

func (c Account) Register() revel.Result {
  return c.Render()
}

func (c Account) HandleRegister(user *models.UserAccount) revel.Result {
  user.Validate(c.Validation)

  // Handle errors
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()    
    return c.Redirect(Account.Register)
  }

  // Ok, display the created user
  return c.Redirect(App.Index)
}