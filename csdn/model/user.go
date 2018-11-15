package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*

create database test;

create table user_info(
	id int(11) auto_increment primary key,
	username varchar(64) not null default ''
);

*/
type Person struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}

func (p *Person) FetchSingleUser(id int) {
	db, err := sql.Open("mysql", "test:123456@/test?charset=utf8")

	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)

	row := db.QueryRow("select id, username from user_info where id = ?", id)
	err = row.Scan(&p.Id, &p.UserName)
	checkErr(err)
}


















