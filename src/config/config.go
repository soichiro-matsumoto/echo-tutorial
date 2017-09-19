package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DbConfig struct {
	DBMS     string
	USER     string
	PASS     string
	PROTOCOL string
	DBNAME   string
}

func (config *DbConfig) GormOpen() (db *gorm.DB, err error) {
	db, err = gorm.Open(config.DBMS, (config.USER + ":" + config.PASS + "@" + config.PROTOCOL + "/" + config.DBNAME + "?charset=utf8&parseTime=True&loc=Local"))

	if err != nil {
		return
	}

	db.LogMode(true)

	return
}

var Config = &DbConfig{
	DBMS:     "mysql",
	USER:     "root",
	PASS:     "abcd1234",
	PROTOCOL: "tcp(localhost:3306)",
	DBNAME:   "sakila",
}
