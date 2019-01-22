package main

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDBConnection(t *testing.T){
	//establish connection
	db:=newDb()
	defer db.Close()

	t.Run("CreateTable", func(t *testing.T) {
		createTable(db)
	})

	
	t.Run("InsertIntoTable", func(t *testing.T) {
		insertIntoTable(db)
	})
	
	//drop table

}

