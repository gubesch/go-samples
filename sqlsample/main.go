package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	ID   int    `json:"ID"`
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

type Item struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	QualityLevel string `json:"QualityLevel"`
	User *User `json:"user"`
} 

func newDb() sql.DB {

	//function for creating a database connection

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

func createTable(db sql.DB) error {

	//Create create statements for user and item table

	createUser:= `
		CREATE TABLE IF NOT EXISTS user(
			ID int NOT NULL auto_increment, 
			firstname varchar(100), 
			lastname varchar(100), PRIMARY KEY (ID)
		) ;`
	createItem:= `
		CREATE TABLE IF NOT EXISTS item (
    		ID int AUTO_INCREMENT,
    		Name VARCHAR(100) NOT NULL,
     		QualityLevel int not null,
			UserID int  NOT NULL ,
    		PRIMARY KEY (ID),
    		CONSTRAINT FK_UserTable
    		FOREIGN KEY (UserID) REFERENCES user(ID)
		);`

	//execute these commands on the database

	_,err := db.Exec(createUser)
	if err != nil{
		return err
	}
	_,err = db.Exec(createItem)
	if err != nil{
		return err
	}
	return nil
}

func insertIntoTable(db sql.DB) error{
	//create prepared statement for user table
	stmtUser,err := db.Prepare("INSERT INTO `user` (`firstname`,`lastname`) VALUES (?, ?)")
	defer stmtUser.Close()
	if err != nil {
		return err
	}

	//execute these prepared statements on the db
	_,err = stmtUser.Exec("Anton", "Horvath")
	if err != nil {
		return err
	}
	_,err =stmtUser.Exec("David", "Kröll")
	if err != nil {
		return err
	}
	_,err =stmtUser.Exec("Harry", "G")
	if err != nil {
		return err
	}

	stmtItem,err := db.Prepare("INSERT INTO `item` (`Name`,`QualityLevel`, `UserID`) VALUES (?, ?, ?)")
	defer stmtItem.Close()
	if err != nil {
		return err
	}

	_,err = stmtItem.Exec("husqvarna fichtenmoped" , 10, 2)
	if err != nil {
		return err
	}
	_,err = stmtItem.Exec("makita gerät" , 10, 1)
	if err != nil {
		return err
	}
	_,err = stmtItem.Exec("Weißbier" , 10, 3)
	if err != nil {
		return err
	}


	return nil

}

//function to truncate a table

func truncateTable(db sql.DB) error {
	_,err := db.Exec("TRUNCATE TABLE user;")
	if err != nil {
		return err
	}
	return nil
}

func selectTableData(db sql.DB) error {
	//create select statement for user table
	usersQuery, err := db.Query("SELECT * FROM `user`;")
	defer usersQuery.Close()
	if err != nil {
		return nil
	}

	//foreach row from the statement above
	for usersQuery.Next(){
		var user User
		err = usersQuery.Scan(&user.ID,&user.firstname,&user.lastname)
		if err != nil{
			return err
		}
		//create the query for the items belonging to specific user
		itemQuery,err := db.Query("SELECT ID, Name, QualityLevel FROM item WHERE UserID = ?",user.ID)
		if err != nil{
			return err
		}
		//foreach item create the item object with the user object
		for itemQuery.Next() {
			var item Item
			err = itemQuery.Scan(&item.ID,&item.Name,&item.QualityLevel)
			if err != nil{
				return err
			}
			item.User = &user
			fmt.Println(item.ID,item.Name,item.QualityLevel,item.User)
		}
	}
	return nil
}


