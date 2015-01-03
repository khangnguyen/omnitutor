package controllers

import (
  "database/sql"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/revel/revel"
)

type GormController struct {
  *revel.Controller
  Tx *gorm.DB
}

var Db gorm.DB

func InitDB() {
  var err error
  ConnectionString, _ := revel.Config.String("db.connectionstring")
  Db, err = gorm.Open("postgres", ConnectionString)
  if err != nil {
    revel.ERROR.Println("FATAL", err)
    panic(err)
  }
}

func (c *GormController) Begin() revel.Result {
  txn := Db.Begin()
  if txn.Error != nil {
    panic(txn.Error)
  }
  c.Tx = txn
  revel.INFO.Println("c.Tx init", c.Tx)
  return nil
}

func (c *GormController) Commit() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Commit()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  revel.INFO.Println("c.Tx commited (nil)")
  return nil
}

func (c *GormController) Rollback() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Rollback()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  return nil
}