package main

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDBConnection(t *testing.T){
	//establish connection
	db:=newDb()
	defer db.Close()

	t.Run("CreateTables", func(t *testing.T) {
		err :=createTable(db)
		if err != nil {
			t.Errorf(err.Error())
		}
	})


	
	t.Run("InsertIntoTables", func(t *testing.T) {
		err := insertIntoTable(db)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("SelectFromTables", func(t *testing.T) {
		err := selectTableData(db)
		if err != nil{
			t.Errorf(err.Error())
		}
	})
	/*
	t.Run("TruncateTable", func(t *testing.T) {
		err := truncateTable(db)
		if err != nil {
			t.Errorf(err.Error())
		}
	})*/


}

