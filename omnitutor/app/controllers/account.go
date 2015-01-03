package controllers

import (
  "github.com/revel/revel"
)

type Account struct {
  *revel.Controller
}

func (c Account) Register() revel.Result {
  return c.Render()
}