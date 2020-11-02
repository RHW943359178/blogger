package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init(dns string) (err error) {
	db, err = sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Println("connect to database failed, err: ", err)
		return
	}
	//	查看是否连接成功
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
