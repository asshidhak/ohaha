package app

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
var Db *sql.DB

func init() {
	cfgBytes, err := ioutil.ReadFile("app.json")
	if err != nil {
		panic(err)
		return
	}

	if err := json.Unmarshal(cfgBytes, &Cfg); err != nil {
		panic(err)
		return
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		Cfg.Mysql.User,
		Cfg.Mysql.Password,
		Cfg.Mysql.Host,
		Cfg.Mysql.Port,
		Cfg.Mysql.Db,
	)

	db := &sql.DB{}
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		//panic(err)
		//return
	}
	if err = db.Ping(); err != nil {
		//panic(err)
		//return
	}

	Db = db
	//defer db.Close()

}