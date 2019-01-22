package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)



func newDb() sql.DB {
	driver := "mysql"
	username := "root"
	password := ""
	domain := "(127.0.0.1:3306)"
	dbName := "gouser"

	database, err := sql.Open(driver, username + ":" + password + "@" + domain + "/" + dbName)
	if err != nil {
		log.Fatal(err)
	}
	return *database;
}

func main(){
	fmt.Printf("SQL Sample")

	db:=newDb()
	defer db.Close()

	createTable(db)
	insertIntoTable(db)
}

func createTable(db sql.DB) {
	_,err := db.Exec("CREATE TABLE IF NOT EXISTS `user`(`ID` int(11) unsigned NOT NULL auto_increment, `firstname` varchar(100), `lastname` varchar(100), PRIMARY KEY (ID)) ;")
	if err != nil{
		log.Println(err)
	}
}

func insertIntoTable(db sql.DB){

	stmt,err := db.Prepare("INSERT INTO `user` (`firstname`,`lastname`) VALUES (?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}

	stmt.Exec("Anton", "Horvath")
	stmt.Exec("David", "Kröll")
	stmt.Exec("Harry", "G")



}
