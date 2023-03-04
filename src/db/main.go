package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/testDB")
	if err != nil {
		fmt.Println(err.Error(), "1111111A")
		return
	}
	defer db.Close()
	/*
		creat table is not exists 'users'(
			'id' int(100) not null auto_increment,
			'username' varchar(45) not null,
			'password' varchar(40) not null,
			'isActive' tinyint(1) default null,
			primary key ('id'),
			unique key 'id_unique' ('id')
		)
	*/
	// query := `creat table is not exists 'user'(
	// 	'id' int(100) not null auto_increment,
	// 	'username' varchar(45) not null,
	// 	'password' varchar(40) not null,
	// 	'isActive' tinyint(1) default null,
	// 	primary key ('id'),
	// 	unique key 'id_unique' ('id')
	// )`
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS mytable (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, some_text TEXT NOT NULL);")

	if err != nil {
		fmt.Println(err.Error(), "2222222A")
		return
	}
	fmt.Println("blah blah ")
}
