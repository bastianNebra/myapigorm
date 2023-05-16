package main

import (
	_ "database/sql"
	"fmt"
	_"log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password="1234"
	dbname = "mydb"
	
)
var myLogger = NewLogger()
//Struct Servershop
type ServerShop struct {
	DB *gorm.DB
}

// Open a database connexion
func (sshop *ServerShop) openDBCon() error{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=disable",host,port,user,password,dbname)

	//DB connexion with GORM
	db, err := gorm.Open(postgres.Open(psqlInfo),&gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("failed to connect database: %v", err)
		myLogger.ErrorLoggerFunc(msg)
	}

	sshop.DB = db

	// Migrate the schema
	err = db.AutoMigrate(&Customer{},&Product{})
	if err !=nil {
		msg := fmt.Sprintf("failed to connect database: %v", err)
		myLogger.ErrorLoggerFunc(msg)
	}
	

	return nil
}
